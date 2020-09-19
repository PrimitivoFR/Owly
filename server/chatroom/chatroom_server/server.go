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
	user_ids := req.GetUsers()
	if name == "" {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("The chatroom name can't be empty"))
	}
	chatroom := models.Chatroom{
		ID:    primitive.NewObjectID(),
		Name:  name,
		Users: user_ids,
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
	for _, user_id := range user_ids {

		filter := bson.M{"_id": user_id}
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
	filter := bson.M{"_id": user_id}
	var user_result models.UserMongo

	err := common_mongo.UserCollection.FindOne(context.Background(), filter).Decode(&user_result)

	if err != nil { //err typically is "no documents in result"
		log.Printf("Error while getting chatrooms for user %v. Error is: %v", user_result.Username, err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while getting chatrooms for user %v. Error is: %v", user_result.Username, err),
		)
	}

	// build the list of chatroom objects
	users_chatrooms := []*chatroompb.Chatroom{}

	for _, chatroom_oid := range user_result.Chatrooms {

		oid, err_oid := primitive.ObjectIDFromHex(chatroom_oid)

		if err_oid != nil {
			log.Printf("Error while gathering chatrooms (oid). Chatroom: %v. Error: %v", chatroom_oid, err_oid)
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while gathering chatrooms (oid). Chatroom: %v. Error: %v", chatroom_oid, err_oid),
			)
		}

		var chatroom_result models.Chatroom
		chatroom_err := common_mongo.ChatroomCollection.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&chatroom_result)

		if chatroom_err != nil {
			log.Printf("Error while gathering chatrooms. Chatroom: %v. Error: %v", chatroom_oid, chatroom_err)
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while gathering chatrooms. Chatroom: %v. Error: %v", chatroom_oid, chatroom_err),
			)
		}

		// TODO : we may also want the chatroom ID ?
		users_chatrooms = append(
			users_chatrooms,
			&chatroompb.Chatroom{
				Name: chatroom_result.Name,
				Id:   chatroom_result.ID.Hex(),
			},
		)
	}

	return &chatroompb.GetChatroomsByUserResponse{
		Chatrooms: users_chatrooms,
		Success:   true,
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
