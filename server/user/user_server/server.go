package user_server

import (
	"context"
	"fmt"
	"log"
	"net"

	"primitivofr/owly/server/common/interceptors"
	common_keycloak "primitivofr/owly/server/common/keycloak"

	"primitivofr/owly/server/user/userpb"

	common_jwt "primitivofr/owly/server/common/jwt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct{}

//

func (*server) SearchUserByUsername(ctx context.Context, req *userpb.SearchUserByUsernameRequest) (*userpb.SearchUserByUsernameResponse, error) {

	currentUserID, err := common_jwt.ReadUUIDFromContext(ctx)
	if err != nil {
		log.Printf("An error has occured while trying to read the token %v", err)
		return nil, status.Error(codes.Internal, fmt.Sprintf("An error has occured while trying to read the token %v", err))

	}

	adminGuy, err := common_keycloak.InitAdmin()
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
		if userRes.Uuid != currentUserID {
			userListRes = append(userListRes, userRes)
		}
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

	s := grpc.NewServer(grpc.UnaryInterceptor(interceptors.UnaryInterceptor))
	userpb.RegisterUserServiceServer(s, &server{})
	reflection.Register(s) // allows us to expose the gRPC server so the client can see what's available. You can use Evans CLI for that

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
