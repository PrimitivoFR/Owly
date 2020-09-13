package message_server

import (
	"encoding/json"
	"log"
	"net"
	message_models "primitivofr/owly/server/message/message_server/models"
	"primitivofr/owly/server/message/messagepb"
	"time"

	"github.com/rgamba/evtwebsocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
