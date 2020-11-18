package chatroom_server

import (
	"context"
	"fmt"
	"log"
	"net"
	"primitivofr/owly/server/chatroom/chatroompb"
	"primitivofr/owly/server/common/interceptors"
	common_jwt "primitivofr/owly/server/common/jwt"
	common_keycloak "primitivofr/owly/server/common/keycloak"
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
	currentUserID, err := common_jwt.ReadUUIDFromContext(ctx)

	log.Println(currentUserID)

	if err != nil {
		return nil, err
	}
	name := req.GetName()
	user_ids := req.GetUsers()

	if name == "" {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("The chatroom name can't be empty"))
	}

	if len([]rune(name)) > 50 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Chatroom name is too long. Req: %v", req))
	}

	user_ids = append(user_ids, currentUserID)

	// getting rid of duplicated users in user_ids
	var deduplicated_user_ids []string
	user_map := make(map[string]bool)

	for _, user := range user_ids {
		if !user_map[user] {
			deduplicated_user_ids = append(deduplicated_user_ids, user)
			user_map[user] = true
		}
	}
	user_ids = deduplicated_user_ids

	chatroom := models.Chatroom{
		ID:    primitive.NewObjectID(),
		Owner: currentUserID,
		Name:  name,
		Users: user_ids,
	}

	res, err := common_mongo.InsertOneChatroomCollection(chatroom)

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
	// user_id := req.GetUserID()
	user_id, err := common_jwt.ReadUUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user_result, err := common_mongo.FindOneUserCollection(user_id)

	if err != nil { //err typically is "no documents in result"
		log.Printf("Error while getting chatrooms for user %v. Error is: %v", user_result.Username, err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while getting chatrooms for user %v. Error is: %v", user_result.Username, err),
		)
	}

	// build the list of chatroom objects
	users_chatrooms := []*chatroompb.Chatroom{}

	for _, chatroomOid := range user_result.Chatrooms {

		// No need to convert to object ID, it's done in the function, since we know we use
		// objectID as value for _id for a chatroom object in chatroom collection
		chatroomResult, chatroomErr := common_mongo.FindOneChatroomCollection(chatroomOid)

		if chatroomErr != nil {
			log.Printf("Error while gathering chatrooms. Chatroom: %v. Error: %v", chatroomOid, chatroomErr)
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while gathering chatrooms. Chatroom: %v. Error: %v", chatroomOid, chatroomErr),
			)
		}

		// Retrieve usernames from userIds
		adminGuy, err := common_keycloak.InitAdmin()
		var usersList []*chatroompb.ChatroomUser

		if err != nil {
			log.Printf("Error while intialiazing connection to keycloak. Error: %v", err)
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while intialiazing connection to keycloak. Error: %v", err),
			)
		}

		for _, uuid := range chatroomResult.Users {
			res, err := adminGuy.GetUserByUUID(uuid)
			chatroomUser := &chatroompb.ChatroomUser{Uuid: uuid}
			if err == nil {
				username := *res.Username
				chatroomUser.Username = username
			} else {
				chatroomUser.Username = "Unknown"
			}

			usersList = append(usersList, chatroomUser)
		}

		//

		// TODO : we may also want the chatroom ID ?
		users_chatrooms = append(
			users_chatrooms,
			&chatroompb.Chatroom{
				Name:  chatroomResult.Name,
				Id:    chatroomResult.ID.Hex(),
				Owner: chatroomResult.Owner,
				Users: usersList,
			},
		)
	}

	return &chatroompb.GetChatroomsByUserResponse{
		Chatrooms: users_chatrooms,
		Count:     int64(len(users_chatrooms)),
		Success:   true,
	}, nil
}

func (*server) DeleteChatroom(ctx context.Context, req *chatroompb.DeleteChatroomRequest) (*chatroompb.DeleteChatroomResponse, error) {

	// get ID of current user
	currentUserID, err := common_jwt.ReadUUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	targetChatroomID := req.ChatroomId

	// check if the current user is the rightful owner
	currentUserIsOwner, err := common_mongo.IsChatroomOwner(currentUserID, targetChatroomID)
	if err != nil {
		return nil, err
	}

	if !currentUserIsOwner {
		return nil, status.Error(
			codes.PermissionDenied,
			fmt.Sprintf(
				"Current user is not the rightful owner of this chatroom, and thus cannot delete it."),
		)
	}

	// get targetChatroom from targetChatroomID
	targetChatroom, errGetChatroom := common_mongo.FindOneChatroomCollection(targetChatroomID)
	if errGetChatroom != nil {
		return nil, errGetChatroom
	}

	// make all users leave the targetChatroom
	for _, userID := range targetChatroom.Users {
		common_mongo.PopChatroomInUserCollection(userID, targetChatroomID)
	}

	// delete chatroom object from database
	deleteErr := common_mongo.DeleteOneChatroomCollection(targetChatroomID)
	if deleteErr != nil {
		return nil, deleteErr
	}

	return &chatroompb.DeleteChatroomResponse{Success: true}, nil
}

func (*server) LeaveChatroom(ctx context.Context, req *chatroompb.LeaveChatroomRequest) (*chatroompb.LeaveChatroomResponse, error) {

	chatroomID := req.ChatroomId
	userID, err := common_jwt.ReadUUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// check if the current user is the rightful owner
	userIsOwner, ownerErr := common_mongo.IsChatroomOwner(userID, chatroomID)
	if ownerErr != nil {
		return nil, ownerErr
	}

	if userIsOwner {
		return nil, status.Error(
			codes.PermissionDenied,
			fmt.Sprintf("Current user is the rightful owner of this chatroom, and thus cannot leave it."),
		)
	}

	// pop user from the user list in the chatroom object
	// and vice versa
	popUserChatroomErr := common_mongo.PopUserInChatroomCollection(userID, chatroomID)
	popChatroomUserErr := common_mongo.PopChatroomInUserCollection(userID, chatroomID)

	if popUserChatroomErr != nil {
		return nil, popUserChatroomErr
	}
	if popChatroomUserErr != nil {
		return nil, popChatroomUserErr
	}

	return &chatroompb.LeaveChatroomResponse{Success: true}, nil
}

func (*server) TransferOwnership(ctx context.Context, req *chatroompb.TranferOwnershipRequest) (*chatroompb.TranferOwnershipResponse, error) {

	userId, err := common_jwt.ReadUUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// check if the current user is the rightful owner
	userIsOwner, ownerErr := common_mongo.IsChatroomOwner(userId, req.ChatroomId)
	if ownerErr != nil {
		return nil, ownerErr
	}

	if !userIsOwner {
		return nil, status.Error(
			codes.PermissionDenied,
			fmt.Sprintf("Current user is not the rightful owner of this chatroom, and thus cannot transfer its ownership."),
		)
	}

	if userId == req.NewOwnerId {
		return nil, status.Error(
			codes.InvalidArgument,
			fmt.Sprintf(
				"It's not really useful to transfer ownership to yourself."),
		)
	}

	targetChatroom, chatroomErr := common_mongo.FindOneChatroomCollection(req.ChatroomId)
	if chatroomErr != nil {
		return nil, chatroomErr
	}

	// check that the new owner is in chatroom
	newOwnerInChatroom := false

	for _, userID := range targetChatroom.Users {
		if userID == req.NewOwnerId {
			newOwnerInChatroom = true
			break
		}
	}

	if !newOwnerInChatroom {
		return nil, status.Error(
			codes.PermissionDenied,
			fmt.Sprintf("New owner is not found in the chatroom"),
		)
	}

	changeOwnerErr := common_mongo.ChangeChatroomOwner(req.NewOwnerId, req.ChatroomId)
	if changeOwnerErr != nil {
		return nil, changeOwnerErr
	}

	return &chatroompb.TranferOwnershipResponse{Success: true}, nil
}

//

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

	s := grpc.NewServer(grpc.UnaryInterceptor(interceptors.UnaryInterceptor))
	chatroompb.RegisterChatroomServiceServer(s, &server{})
	reflection.Register(s) // allows us to expose the gRPC server so the client can see what's available. You can use Evans CLI for that

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
