package authserver

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"primitivofr/owly/server/auth/authpb"
	"reflect"
	"testing"

	kvdb "primitivofr/owly/server/common/kvdb"

	"github.com/boltdb/bolt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var db *bolt.DB

func init() {
	var err error
	db, err = kvdb.InitDB("testingdb", "common")
	if err != nil {
		panic(err)
	}

}

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

		resp, err := s.CreateNewUser(context.Background(), req)

		assert.Nil(t, err, "CreateNewUser got unexpected error")
		if o := assert.Equal(t, resp.Success, tt.want.Success, "Created user does not correspond to the wanted user"); o {
			b, err := json.Marshal(tests[0].req.Username)
			if err != nil {
				panic(err)
			}

			err = kvdb.WriteData(db, "users", tests[0].req.Username, string(b))
			if err != nil {
				panic(err)
			}
		}
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

			err := kvdb.WriteData(db, "tokens", "applinh", resp.Result.AccessToken)
			if err != nil {
				panic(err)
			}

			// This needs to be removed
			f, err := os.Create("../../token.txt")
			if err != nil {
				panic(err)
			}
			f.WriteString(resp.Result.GetAccessToken())

		}
	}
}
