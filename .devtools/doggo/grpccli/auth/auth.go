package auth

import (
	"context"

	"github.com/primitivofr/owly/server/common/pb/auth/authpb"

	"google.golang.org/grpc"
)

// LoginUser loginUser grpc call
func LoginUser(username string, password string) *authpb.LoginUserResponse {
	conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	c := authpb.NewAuthServiceClient(conn)
	req := &authpb.LoginUserRequest{
		Username: username,
		Password: password,
	}
	res, err := c.LoginUser(context.Background(), req)
	if err != nil {
		panic(err)
	}

	return res

}
