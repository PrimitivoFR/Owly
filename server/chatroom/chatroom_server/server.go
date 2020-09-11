package chatroom_server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"primitivofr/owly/server/chatroom/chatroom_server/models"
	"primitivofr/owly/server/chatroom/chatroompb"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	res, err := chatroomCollection.InsertOne(context.Background(), chatroom)
	if err != nil {
		log.Fatalf("Error while creating chatroom: %v. Error is: %v", chatroom, err)
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

	return &chatroompb.GetChatroomsByUserResponse {
		Chatrooms: nil,
		Success: false, // TODO : STUB
	}, nil
}

//
// Global vars
var chatroomCollection *mongo.Collection
var userCollection *mongo.Collection
var client *mongo.Client

func setupMongoDB() error {
	var err error
	fmt.Println("Connecting to MongoDB...")
	MONGO_HOSTNAME := "localhost"
	MONGO_HOSTNAME = os.Getenv("MONGO_HOSTNAME")
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://" + MONGO_HOSTNAME + ":27017"))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	chatroomCollection = client.Database("owly").Collection("chatrooms")
	userCollection = client.Database("owly").Collection("users")
	return nil
}

func StartServer() {

	lis, err := net.Listen("tcp", "0.0.0.0:50052")

	if err != nil {
		log.Fatalf("Failed to listen %v \n", err)
	}

	// -----MONGO-------

	err = setupMongoDB()
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
