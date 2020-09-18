package chatroom_server

import (
	"context"
	"fmt"
	"log"
	"net"
	"primitivofr/owly/server/chatroom/chatroompb"
	"primitivofr/owly/server/common/models"
	common_mongo "primitivofr/owly/server/common/mongo"

	"go.mongodb.org/mongo-driver/bson"
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

	// add chatroom to users
	// _, err_update := common_mongo.UserCollection.InsertOne(context.Background(), user_mongo)
	for _, username := range users {

		// check if the user exists and find its id
		user_filter := bson.M{"username": bson.M{"$eq": username}}
		var user_result models.User
		err_user := common_mongo.UserCollection.FindOne(context.Background(), user_filter).Decode(&user_result)

		if err_user != nil {
			log.Printf("Error while creating chatroom (user not found): %v. User: %v", chatroom, username)
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while creating chatroom (user not found): %v. User: %v", chatroom, username),
			)
		}

		// update the chatroom list of the user
		filter := bson.M{"_id": bson.M{"$eq": user_result.ID}}
		update := bson.M{"$push": bson.M{"chatrooms": oid}}
		_, err_update := common_mongo.UserCollection.UpdateOne(context.Background(), filter, update)

		if err_update != nil {
			log.Printf("Error while creating chatroom (adding chatroom to user): %v. Error is: %v", chatroom, err_update)
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while creating chatroom (adding chatroom to user): %v. Error is: %v", chatroom, err_update),
			)
		}
	}

	return &chatroompb.CreateChatroomResponse{
		ID:      oid.Hex(),
		Success: true,
	}, nil

}

func (*server) GetChatroomsByUser(ctx context.Context, req *chatroompb.GetChatroomsByUserRequest) (*chatroompb.GetChatroomsByUserResponse, error) {
	user_id := req.GetUserID()
	filter := bson.M{"_id": bson.M{"$eq": user_id}}
	var user_result models.User

	err := common_mongo.UserCollection.FindOne(context.Background(), filter).Decode(&user_result)

	if err != nil { //err typically is "no documents in result"
		log.Printf("Error while getting chatrooms for user %v. Error is: %v", user_result.Username, err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while getting chatrooms for user %v. Error is: %v", user_result.Username, err),
		)
	}

	//TODO return chatroom[] object
	log.Printf("%v's chatrooms : %v", user_result.Username, user_result.Chatrooms) // todo delete

	return &chatroompb.GetChatroomsByUserResponse{
		Chatrooms: nil,   // TODO : STUB
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
