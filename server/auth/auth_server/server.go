package authserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"primitivofr/owly/server/auth/authpb"
	common_keycloak "primitivofr/owly/server/common/keycloak"
	"primitivofr/owly/server/common/models"
	common_mongo "primitivofr/owly/server/common/mongo"

	"github.com/Nerzal/gocloak/v7"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct{}

//
func (*server) CreateNewUser(ctx context.Context, req *authpb.CreateNewUserRequest) (*authpb.CreateNewUserResponse, error) {

	adminGuy, err := common_keycloak.InitAdmin()
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
	userMongo := models.UserMongo{
		// ID:       primitive.NewObjectID(),
		ID:        ID,
		Username:  username,
		Chatrooms: []string{}, //empty string array
	}

	errInsert := common_mongo.InsertOneUserCollection(userMongo)

	if errInsert != nil {
		log.Printf("Error while creating user: %v. Error is: %v", userMongo, errInsert)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while creating user: %v. Error is: %v", userMongo, errInsert),
		)
	}

	res := &authpb.CreateNewUserResponse{
		Success: true,
	}

	return res, nil
}

func (*server) LoginUser(ctx context.Context, req *authpb.LoginUserRequest) (*authpb.LoginUserResponse, error) {
	adminGuy, err := common_keycloak.InitAdmin()
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

	return &authpb.LoginUserResponse{
		Result: &authpb.JWT{
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

//
func StartServer() {

	lis, err := net.Listen("tcp", "0.0.0.0:50054")

	if err != nil {
		log.Fatalf("Failed to listen %v \n", err)
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &server{})

	reflection.Register(s) // allows us to expose the gRPC server so the client can see what's available. You can use Evans CLI for that

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
