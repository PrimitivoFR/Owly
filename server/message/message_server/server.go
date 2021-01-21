package message_server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	common_elastic "github.com/primitivofr/owly/server/common/elastic"
	common_interceptors "github.com/primitivofr/owly/server/common/interceptors"
	common_jwt "github.com/primitivofr/owly/server/common/jwt"
	"github.com/primitivofr/owly/server/message/interceptors"
	"github.com/primitivofr/owly/server/message/messagepb"
	message_models "github.com/primitivofr/owly/server/message/models"

	"google.golang.org/grpc/codes"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/olivere/elastic"
	"github.com/rgamba/evtwebsocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct{}

//

func (*server) StreamMessagesByChatroom(req *messagepb.StreamMessagesByChatroomRequest, stream messagepb.MessageService_StreamMessagesByChatroomServer) error {
	chatroomID := req.GetChatroomID()

	userID, err := common_jwt.ReadUUIDFromContext(stream.Context())
	if err != nil {
		return err
	}

	err = common_interceptors.IsUserInChatroom(chatroomID, userID)
	if err != nil {
		return err
	}

	var wsError error

	c := evtwebsocket.Conn{
		Reconnect:        true,
		PingIntervalSecs: 5,

		OnConnected: func(w *evtwebsocket.Conn) {
			log.Println("Connected to ES WS")
		},
		OnMessage: func(msg []byte, w *evtwebsocket.Conn) {

			var esWSResp message_models.ESWebSocketResp
			err := json.Unmarshal(msg, &esWSResp)
			if err != nil {
				log.Printf("Error while reading message %v", err)
			} else if esWSResp.Index == chatroomID {
				// We basically listen to ALL entries in the ES db, so we have to filter out to listen to a specific chatroomID
				// However, it would be better to directly listen on a URI like ws://elasticsearch:9400/<chatroomID>/ws/_changes
				// Should investigate that later

				var message = esWSResp.Source
				message.Id = esWSResp.ID

				if err != nil {
					log.Printf("Error while reading message %v", err)
				} else {
					err := stream.Send(&messagepb.StreamMessagesByChatroomResponse{
						Operation: esWSResp.Operation,
						Message:   &message,
					})
					if err != nil {
						log.Printf("Error while streaming %v", err)
					}
				}
			}

		},
		OnError: func(err error) {
			wsError = err
			log.Printf("Error while listening to WS %v", err)
		},
	}
	// Connect
	c.Dial("ws://elasticsearch-srv:9400/ws/_changes", "")

	for c.IsConnected() {
		time.Sleep(3 * time.Second)
	}

	return wsError
}
func (*server) SendMessage(ctx context.Context, req *messagepb.SendMessageRequest) (*messagepb.SendMessageResponse, error) {

	user_id, err := common_jwt.ReadUUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if len([]rune(req.Message.Content)) > 1000 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Message is too long. Req: %v", req))
	}

	err = common_interceptors.IsUserInChatroom(req.Message.ChatroomID, user_id)
	if err != nil {
		return nil, err
	}

	client, err := common_elastic.Init()
	if err != nil {
		return nil, err
	}

	// Create a mapping for the Elasticsearch documents
	var docMap map[string]interface{}

	req.Message.AuthorUUID = user_id
	req.Message.History = []*messagepb.MessageHistory{} // Init history as empty array
	// TODO: Read authorNAME too from token

	req.Message.Timestamp = int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
	doc, err := json.Marshal(req.Message)
	if err != nil {
		log.Printf("Json Marshaling returned error : %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Json Marshaling returned error: %v", err),
		)
	}

	indexRequest := esapi.IndexRequest{
		Index:   req.Message.ChatroomID,
		Body:    strings.NewReader(string(doc)),
		Refresh: "true",
	}

	res, err := indexRequest.Do(ctx, client)
	if err != nil {
		log.Printf("Inserting index failed: %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Inserting index failed: %v", err),
		)
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Inserting index failed: %v", res.Status())
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Inserting index failed: %v", err),
		)
	}
	if err := json.NewDecoder(res.Body).Decode(&docMap); err != nil {
		log.Printf("Error parsing the response body: %s", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error parsing the response body: %v", err),
		)
	}

	msgID := docMap["_id"]

	response := messagepb.SendMessageResponse{
		Success: true,
		Id:      msgID.(string),
	}
	return &response, nil
}

func (*server) GetMessagesByChatroom(ctx context.Context, req *messagepb.GetMessagesByChatroomRequest) (*messagepb.GetMessagesByChatroomResponse, error) {

	user_id, err := common_jwt.ReadUUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	err = common_interceptors.IsUserInChatroom(req.ChatroomID, user_id)
	if err != nil {
		return nil, err
	}

	client, err := common_elastic.Init()
	if err != nil {
		return nil, err
	}

	res, err := client.Search(
		client.Search.WithIndex(req.GetChatroomID()),
		client.Search.WithSize(10000),
	)

	if err != nil {
		log.Printf("Error while searching in ES: %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while searching in ES: %v", err),
		)
	}

	var messages []*messagepb.Message

	var response map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Printf("Error while parsing body: %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while parsing body: %v", err),
		)
	}

	if _, ok := response["hits"]; !ok {
		return &messagepb.GetMessagesByChatroomResponse{
			Messages: messages,
		}, nil
	}
	var h map[string]interface{} = response["hits"].(map[string]interface{})

	if _, ok := h["hits"]; !ok {
		return &messagepb.GetMessagesByChatroomResponse{
			Messages: messages,
		}, nil
	}

	var hits []interface{} = h["hits"].([]interface{})

	for _, hit := range hits {
		currentHit := hit.(map[string]interface{})
		var message *messagepb.Message

		sourceMap := currentHit["_source"]
		sourceJSON, _ := json.Marshal(sourceMap)
		json.Unmarshal(sourceJSON, &message)

		message.Id = currentHit["_id"].(string)
		messages = append(messages, message)
	}

	return &messagepb.GetMessagesByChatroomResponse{
		Messages: messages,
	}, nil

}

func (*server) DeleteMessage(ctx context.Context, req *messagepb.DeleteMessageRequest) (*messagepb.DeleteMessageResponse, error) {
	userId, err := common_jwt.ReadUUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	err = common_interceptors.IsUserInChatroom(req.ChatroomID, userId)
	if err != nil {
		return nil, err
	}

	//Elastic client init with Olivere package
	client, err := common_elastic.InitOlivereES()
	if err != nil {
		return nil, err
	}

	err = common_elastic.DoesChatroomIndexExist(client, req.ChatroomID)
	if err != nil {
		return nil, err
	}

	// CHECK if userId == authorUUID of the message for which a deletion has been requested
	err = interceptors.IsUserAuthorOfMessage(client, req.ChatroomID, req.MessageID, userId)
	if err != nil {
		return nil, err
	}

	// Delete the message
	deleteService := elastic.NewDeleteService(client)
	deleteService.Index(req.ChatroomID)
	deleteService.Id(req.MessageID)
	deleteService.Type("_doc")
	deleteService.Refresh("true")

	_, err = deleteService.Do(context.Background())

	if err != nil {
		log.Printf("Delete Index returned error %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Delete Index returned error %v", err),
		)
	}

	response := messagepb.DeleteMessageResponse{
		Success: true,
	}

	return &response, nil
}

// UpdateMessageContent updates the content of a message
// it receives messageId as well as the chatroomId which contains the message
// it also receives newContent, which contains the content to update the targeted message with
func (*server) UpdateMessageContent(ctx context.Context, req *messagepb.UpdateMessageContentRequest) (*messagepb.UpdateMessageContentResponse, error) {

	// Retrieve userId from token
	userId, err := common_jwt.ReadUUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// Check if user belongs to chatroom
	err = common_interceptors.IsUserInChatroom(req.ChatroomId, userId)
	if err != nil {
		return nil, err
	}

	//Elastic client init with Olivere package
	client, err := common_elastic.InitOlivereES()
	if err != nil {
		return nil, err
	}

	// Check if index exists
	err = common_elastic.DoesChatroomIndexExist(client, req.ChatroomId)
	if err != nil {
		return nil, err
	}

	// CHECK if userId == authorUUID of the message for which an update has been requested
	err = interceptors.IsUserAuthorOfMessage(client, req.ChatroomId, req.MessageId, userId)
	if err != nil {
		return nil, err
	}

	// Get the message
	getService := elastic.NewGetService(client)
	getService.Index(req.ChatroomId)
	getService.Id(req.MessageId)
	getService.Type("_doc")
	getService.Refresh("true")

	resGet, err := getService.Do(context.Background())
	if err != nil {
		log.Printf("UpdateMessageContent returned error %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("UpdateMessageContent returned error %v", err),
		)
	}

	var message *messagepb.Message

	if err := json.Unmarshal(*resGet.Source, &message); err != nil {
		log.Printf("Error while parsing body: %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while parsing body: %v", err),
		)
	}

	// Update the history
	message.History = append(message.History, &messagepb.MessageHistory{Timestamp: int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond), Content: message.Content})

	// Update the message
	updateService := elastic.NewUpdateService(client)
	updateService.Index(req.ChatroomId)
	updateService.Id(req.MessageId)
	updateService.Type("_doc")
	updateService.Doc(map[string]interface{}{"content": req.NewContent, "history": message.History})
	updateService.Refresh("true")

	res, err := updateService.Do(context.Background())
	if err != nil {
		log.Printf("UpdateMessageContent returned error %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("UpdateMessageContent returned error %v", err),
		)
	}

	return &messagepb.UpdateMessageContentResponse{
		Success: true,
		Message: &messagepb.Message{
			Id: res.Id,
		},
	}, nil
}

// StartServer starts the message grpc server on port 50053
func StartServer() {

	lis, err := net.Listen("tcp", "0.0.0.0:50053")

	if err != nil {
		log.Fatalf("Failed to listen %v \n", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(common_interceptors.UnaryInterceptor), grpc.StreamInterceptor(common_interceptors.StreamInterceptor))
	messagepb.RegisterMessageServiceServer(s, &server{})
	reflection.Register(s) // allows us to expose the gRPC server so the client can see what's available. You can use Evans CLI for that

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
