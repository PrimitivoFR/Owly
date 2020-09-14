package chatroom_server

import (
	"context"
	"fmt"
	"log"
	"net"
	"primitivofr/owly/server/chatroom/chatroompb"
	"primitivofr/owly/server/common/models"
	common_mongo "primitivofr/owly/server/common/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct{}

//
func (*server) CreateChatroom(ctx context.Context, req *chatroompb.CreateChatroomRequest) (*chatroompb.CreateChatroomResponse, error) {

	name := req.GetName()
	users := req.GetUsers()
	if name == "" {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("The chatroom name can't be empty"))
	}
	chatroom := models.Chatroom{
		ID:    primitive.NewObjectID(),
		Name:  name,
		Users: users,
	}

	res, err := common_mongo.ChatroomCollection.InsertOne(context.Background(), chatroom)
	if err != nil {
		log.Printf("Error while creating chatroom: %v. Error is: %v", chatroom, err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while creating chatroom: %v. Error is: %v", chatroom, err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot converted to OID: %v", err),
		)
	}

	return &chatroompb.CreateChatroomResponse{
		ID:      oid.Hex(),
		Success: true,
	}, nil

}

// TODO
func (*server) GetChatroomsByUser(ctx context.Context, req *chatroompb.GetChatroomsByUserRequest) (*chatroompb.GetChatroomsByUserResponse, error) {
	user_name := req.GetUser()

	fmt.Println("DEBUG: got user: ", user_name)

	return &chatroompb.GetChatroomsByUserResponse{
		Chatrooms: nil,
		Success:   false, // TODO : STUB
	}, nil
}

func StartServer() {

	lis, err := net.Listen("tcp", "0.0.0.0:50052")

	if err != nil {
		log.Fatalf("Failed to listen %v \n", err)
	}

	// -----MONGO-------

	err = common_mongo.SetupMongoDB()
	if err != nil {
		log.Fatalf("Erreur setting up MongoDB: %v", err)
		return
	}

	// -----MONGO-------

	s := grpc.NewServer()
	chatroompb.RegisterChatroomServiceServer(s, &server{})
	reflection.Register(s) // allows us to expose the gRPC server so the client can see what's available. You can use Evans CLI for that

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
