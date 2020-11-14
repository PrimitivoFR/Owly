package authserver

import (
	"context"
	"fmt"
	"log"
	"os"
	"primitivofr/owly/server/auth/authpb"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateNewUser(t *testing.T) {

	s := server{}

	tests := []struct {
		req  authpb.CreateNewUserRequest
		want authpb.CreateNewUserResponse
	}{
		{
			req: authpb.CreateNewUserRequest{
				Email:     "toto@toto.fr",
				Password:  "Aze123",
				Username:  "applinh",
				FirstName: "Thomas",
				LastName:  "Martin",
			},
			want: authpb.CreateNewUserResponse{
				Success: true,
			},
		},
	}

	for _, tt := range tests {
		req := &tt.req

		log.Println(req)

		resp, err := s.CreateNewUser(context.Background(), req)

		assert.Nil(t, err, "CreateNewUser got unexpected error")
		assert.Equal(t, resp.Success, tt.want.Success, "Created user does not correspond to the wanted user")
	}

	log.Println("[Successfully passed TestCreateNewUser]")
}

func TestLoginUser(t *testing.T) {
	s := server{}

	tests := []struct {
		req  authpb.LoginUserRequest
		want interface{}
	}{
		{
			req: authpb.LoginUserRequest{
				Username: "applinh",
				Password: "Aze123",
			},
			want: &authpb.LoginUserResponse{},
		},
		{
			req: authpb.LoginUserRequest{
				Username: "applinh",
				Password: "Aze1234",
			},
			want: status.Error(codes.PermissionDenied, ""),
		},
	}

	for _, tt := range tests {

		resp, err := s.LoginUser(context.Background(), &tt.req)

		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error

			assert.Equal(t, reflect.TypeOf(tt.want), reflect.TypeOf(err), "LoginUser got unexpected error")

		} else {
			assert.Equal(t, reflect.TypeOf(resp), reflect.TypeOf(tt.want))
		}

		if resp != nil {
			f, err := os.Create("../../token.txt")
			if err != nil {
				fmt.Println(err)
				return
			}
			f.WriteString(resp.Result.GetAccessToken())
		}
	}
}
