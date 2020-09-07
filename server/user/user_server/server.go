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

	adminGuy := keycloak.InitAdmin()

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
		log.Fatalf("Error while logging: %v", err)
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