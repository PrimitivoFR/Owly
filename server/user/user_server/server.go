package user_server

import (
	"context"
	"fmt"
	"log"
	"net"
	"primitivofr/owly/server/common/models"
	"primitivofr/owly/server/user/user_server/keycloak"
	"primitivofr/owly/server/user/userpb"

	common_mongo "primitivofr/owly/server/common/mongo"

	"github.com/Nerzal/gocloak/v7"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct{}

//

func (*server) CreateNewUser(ctx context.Context, req *userpb.CreateNewUserRequest) (*userpb.CreateNewUserResponse, error) {

	adminGuy, err := keycloak.InitAdmin()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error: %v", err))
	}

	firstName := req.GetFirstName()
	lastName := req.GetLastName()
	enabled := true
	username := req.GetUsername()
	email := req.GetEmail()
	pass := req.GetPassword()

	user := gocloak.User{
		FirstName: &firstName,
		LastName:  &lastName,
		Email:     &email,
		Enabled:   &enabled,
		Username:  &username,
	}

	if firstName == "" || lastName == "" || email == "" || pass == "" {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Missing core argument. Req: %v", req))
	}

	if len([]rune(firstName)) > 50 ||
		len([]rune(lastName)) > 50 ||
		len([]rune(email)) > 100 ||
		len([]rune(username)) > 50 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("One or more arguments are too long. Req: %v", req))
	}

	ID, err := adminGuy.CreateUser(user)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error while creating user account: %v", err))
	}

	err = adminGuy.SetUserPassword(ID, pass, false)
	if err != nil {
		log.Printf("Error while setting up password: %v", err)
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error while setting up password: %v", err))
	}

	// insert user object in mongodb
	user_mongo := models.UserMongo{
		// ID:       primitive.NewObjectID(),
		ID:        ID,
		Username:  username,
		Chatrooms: []string{}, //empty string array
	}

	// TODO : check what to do with the return, instead of making it "_"
	// TODO: check if UserCollection is nil, and if so, call the common_mongo.SetupMongoDB(). After that, remove the function call from the test
	_, err_insert := common_mongo.UserCollection.InsertOne(context.Background(), user_mongo)
	if err_insert != nil {
		log.Printf("Error while creating user: %v. Error is: %v", user_mongo, err_insert)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while creating user: %v. Error is: %v", user_mongo, err_insert),
		)
	}

	res := &userpb.CreateNewUserResponse{
		Success: true,
	}

	return res, nil
}

func (*server) LoginUser(ctx context.Context, req *userpb.LoginUserRequest) (*userpb.LoginUserResponse, error) {
	adminGuy, err := keycloak.InitAdmin()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error: %v", err))
	}

	client_secret, err := adminGuy.GetClientSecret("owlycli")
	if err != nil {
		log.Printf("Internal error: %v", err)
		return nil, status.Error(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}
	username := req.GetUsername()
	password := req.GetPassword()

	res, err := adminGuy.Client.Login(context.Background(), "owlycli", client_secret, "OWLY", username, password)

	if err != nil {
		log.Printf("Authentification error: %v", err)
		return nil, status.Error(codes.PermissionDenied, fmt.Sprintf("Authentification error: %v", err))
	}

	return &userpb.LoginUserResponse{
		Result: &userpb.JWT{
			AccessToken:      res.AccessToken,
			IDToken:          res.IDToken,
			ExpiresIn:        int64(res.ExpiresIn),
			RefreshExpiresIn: int64(res.RefreshExpiresIn),
			RefreshToken:     res.RefreshToken,
			TokenType:        res.TokenType,
			NotBeforePolicy:  int64(res.NotBeforePolicy),
			SessionState:     res.SessionState,
			Scope:            res.Scope,
		},
	}, nil
}

func (*server) SearchUserByUsername(ctx context.Context, req *userpb.SearchUserByUsernameRequest) (*userpb.SearchUserByUsernameResponse, error) {
	adminGuy, err := keycloak.InitAdmin()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error: %v", err))
	}

	ret, _ := adminGuy.SearchUserByUsername(req.GetUsername())
	var userListRes []*userpb.User

	for _, user := range ret {
		userRes := &userpb.User{
			Uuid:     *user.ID,
			Username: *user.Username,
			Email:    *user.Email,
		}

		userListRes = append(userListRes, userRes)
	}

	return &userpb.SearchUserByUsernameResponse{
		Results: userListRes,
		Count:   int64(len(userListRes)),
	}, nil
}

//
func StartServer() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen %v \n", err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &server{})
	reflection.Register(s) // allows us to expose the gRPC server so the client can see what's available. You can use Evans CLI for that

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
