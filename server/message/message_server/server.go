package message_server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	common_elastic "primitivofr/owly/server/common/elastic"
	"primitivofr/owly/server/common/interceptors"
	common_jwt "primitivofr/owly/server/common/jwt"
	message_models "primitivofr/owly/server/message/message_server/models"
	"primitivofr/owly/server/message/messagepb"
	"strings"
	"time"

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

	user_id, err := common_jwt.ReadUUIDFromContext(stream.Context())
	if err != nil {
		return err
	}

	if ok, err := interceptors.IsUserInChatroom(chatroomID, user_id); !ok || err != nil {
		if !ok && err == nil {
			log.Printf("User %v not found in this chatroom: %v", user_id, chatroomID)
			return status.Errorf(
				codes.Unauthenticated,
				fmt.Sprintf("User %v not found in this chatroom: %v", user_id, chatroomID),
			)
		}

		log.Printf("Error while checking if user belong to chatroom: %v", err)
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while checking if user belong to chatroom: %v", err),
		)

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
	c.Dial("ws://elasticsearch:9400/ws/_changes", "")

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

	if ok, err := interceptors.IsUserInChatroom(req.Message.GetChatroomID(), user_id); !ok || err != nil {
		if !ok && err == nil {
			log.Printf("User %v not found in this chatroom: %v", user_id, req.Message.GetChatroomID())
			return nil, status.Errorf(
				codes.Unauthenticated,
				fmt.Sprintf("User %v not found in this chatroom: %v", user_id, req.Message.GetChatroomID()),
			)
		}

		log.Printf("Error while checking if user belong to chatroom: %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while checking if user belong to chatroom: %v", err),
		)

	}

	client, err := common_elastic.Init()
	if err != nil {
		return nil, err
	}

	// Create a mapping for the Elasticsearch documents
	var docMap map[string]interface{}

	req.Message.AuthorUUID = user_id
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
	response := messagepb.SendMessageResponse{
		Success: true,
	}
	return &response, nil
}

func (*server) GetMessagesByChatroom(ctx context.Context, req *messagepb.GetMessagesByChatroomRequest) (*messagepb.GetMessagesByChatroomResponse, error) {

	user_id, err := common_jwt.ReadUUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if ok, err := interceptors.IsUserInChatroom(req.GetChatroomID(), user_id); !ok || err != nil {
		if !ok && err == nil {
			log.Printf("User %v not found in this chatroom: %v", user_id, req.GetChatroomID())
			return nil, status.Errorf(
				codes.Unauthenticated,
				fmt.Sprintf("User %v not found in this chatroom: %v", user_id, req.GetChatroomID()),
			)
		}

		log.Printf("Error while checking if user belong to chatroom: %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while checking if user belong to chatroom: %v", err),
		)

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

	fmt.Println(response)

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

	if ok, err := interceptors.IsUserInChatroom(req.ChatroomID, userId); !ok || err != nil {
		if !ok && err == nil {
			log.Printf("User %v not found in this chatroom: %v", userId, req.ChatroomID)
			return nil, status.Errorf(
				codes.Unauthenticated,
				fmt.Sprintf("User %v not found in this chatroom: %v", userId, req.ChatroomID),
			)
		}

		log.Printf("Error while checking if user belong to chatroom: %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while checking if user belong to chatroom: %v", err),
		)

	}

	//Elastic client init with Olivere package
	client, err := elastic.NewClient(
		elastic.SetSniff(true),
		elastic.SetURL("http://elasticsearch:9200"),
		elastic.SetHealthcheckInterval(5*time.Second),
	)
	if err != nil {
		log.Printf("Elasticsearch connection error %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Elasticsearch connection error %v", err),
		)
	}

	exist, err := client.IndexExists(req.ChatroomID).Do(context.Background())
	if !exist {
		log.Printf("Index %v not found", req.ChatroomID)
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Index %v not found", req.ChatroomID),
		)
	}

	// CHECK if userId == authorUUID of the message for which a deletion has been requested
	doc, err := client.Get().
		Index(req.ChatroomID).
		Id(req.MessageID).
		Type("_doc").
		Do(context.Background())

	if err != nil {
		log.Printf("Error while getting message in ES %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while getting  message in ES %v", err),
		)
	}
	var message messagepb.Message
	data := *doc.Source
	json.Unmarshal(data, &message)

	// Not authorized
	if message.AuthorUUID != userId {
		log.Printf("This user %v is not the author of the message %v. He can't delete it", userId, message.Id)
		return nil, status.Errorf(
			codes.Unauthenticated,
			fmt.Sprintf("This user %v is not the author of the message %v. He can't delete it", userId, message.Id),
		)
	}
	//

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

//
func StartServer() {

	lis, err := net.Listen("tcp", "0.0.0.0:50053")

	if err != nil {
		log.Fatalf("Failed to listen %v \n", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(interceptors.UnaryInterceptor), grpc.StreamInterceptor(interceptors.StreamInterceptor))
	messagepb.RegisterMessageServiceServer(s, &server{})
	reflection.Register(s) // allows us to expose the gRPC server so the client can see what's available. You can use Evans CLI for that

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
