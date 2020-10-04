package authserver

import (
	"context"
	"fmt"
	"log"
	"os"
	"primitivofr/owly/server/auth/authpb"
	"reflect"
	"testing"

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

		if err != nil {
			t.Errorf("CreateNewUser got unexpected error %v", err)
		} else if resp.Success != tt.want.Success {
			t.Errorf("CreateNewUser(%v)=%v, wanted %v", req, resp, tt.want)
		} else {
			log.Println("[Successfully passed TestCreateNewUser]")
		}

	}
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
		req := &tt.req

		log.Println(req)

		resp, err := s.LoginUser(context.Background(), req)

		if err != nil && reflect.TypeOf(tt.want) != reflect.TypeOf(err) {
			t.Errorf("LoginUser got unexpected error %v", err)
		} else if resp != nil && reflect.TypeOf(resp) != reflect.TypeOf(tt.want) {
			t.Errorf("LoginUser(%v)=%v, wanted %v", req, resp, reflect.TypeOf(tt.want).String())
		} else {
			log.Println("[Successfully passed TestLoginUser]")

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

}
