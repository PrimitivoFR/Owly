package message_server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	message_models "primitivofr/owly/server/message/message_server/models"
	"primitivofr/owly/server/message/messagepb"
	"strings"
	"time"

	"google.golang.org/grpc/codes"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/rgamba/evtwebsocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct{}

//

func (*server) StreamMessagesByChatroom(req *messagepb.StreamMessagesByChatroomRequest, stream messagepb.MessageService_StreamMessagesByChatroomServer) error {
	chatroomID := req.GetChatroomID()

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
						Message: &message,
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

	// Create a mapping for the Elasticsearch documents
	var docMap map[string]interface{}

	// Declare an Elasticsearch configuration
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://elasticsearch:9200",
		},
	}

	// Connect to ElasticSearch
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Printf("Elasticsearch connection error %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Elasticsearch connection error: %v", err),
		)
	}

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
		log.Printf("Inserting index failed: ", res.Status())
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
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://elasticsearch:9200",
		},
	}

	// Connect to ElasticSearch
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Printf("Elasticsearch connection error %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Elasticsearch connection error: %v", err),
		)
	}
	res, err := client.Search(client.Search.WithIndex(req.GetChatroomID()))

	var response map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Printf("Error while parsing body: %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while parsing body: %v", err),
		)
	}
	var h map[string]interface{} = response["hits"].(map[string]interface{})
	var hits []interface{} = h["hits"].([]interface{})

	var messages []*messagepb.Message

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

//
func StartServer() {

	lis, err := net.Listen("tcp", "0.0.0.0:50053")

	if err != nil {
		log.Fatalf("Failed to listen %v \n", err)
	}

	s := grpc.NewServer()
	messagepb.RegisterMessageServiceServer(s, &server{})
	reflection.Register(s) // allows us to expose the gRPC server so the client can see what's available. You can use Evans CLI for that

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
