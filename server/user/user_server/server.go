package user_server

import (
	"context"
	"fmt"
	"log"
	"net"
	"primitivofr/owly/server/user/user_server/keycloak"
	"primitivofr/owly/server/user/userpb"

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

	ID, err := adminGuy.CreateUser(user)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error while creating user account: %v", err))
	}

	err = adminGuy.SetUserPassword(ID, pass, false)
	if err != nil {
		log.Fatalf("Error while setting up password: %v", err)
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error while setting up password: %v", err))
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

	client_secret, _ := adminGuy.GetClientSecret("owlycli")
	username := req.GetUsername()
	password := req.GetPassword()

	res, err := adminGuy.Client.Login(context.Background(), "owlycli", client_secret, "OWLY", username, password)

	fmt.Println("DEBUG: res: ", res)
	fmt.Println("DEBUG: err: ", err)

	if err != nil {
		log.Fatalf("Authentification error: %v", err)
		return nil, status.Error(codes.PermissionDenied, fmt.Sprintf("Authentification error: %v", err))
	}

	return &userpb.LoginUserResponse{
		Token: &userpb.JWT{
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
