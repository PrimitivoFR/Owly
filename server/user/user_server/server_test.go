package user_server

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"primitivofr/owly/server/user/userpb"
	"reflect"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/stretchr/testify/assert"
)

func check(e error, desc string) {
	if e != nil {
		log.Println(desc)
		panic(e)
	}
}

func TestSearchUserByUsername(t *testing.T) {
	s := server{}

	tests := []struct {
		req  userpb.SearchUserByUsernameRequest
		want userpb.SearchUserByUsernameResponse
	}{
		{
			req: userpb.SearchUserByUsernameRequest{
				Username: "user",
			},
			want: userpb.SearchUserByUsernameResponse{
				Count: 1,
			},
		},
		{
			req: userpb.SearchUserByUsernameRequest{}, // empty username implied
			want: userpb.SearchUserByUsernameResponse{
				// only 1 as this function doesn't return the current user
				// and as "user" should be the only user besides applinh
				Count: 1,
			},
		},
		{
			req: userpb.SearchUserByUsernameRequest{
				Username: "notexist",
			},
			want: userpb.SearchUserByUsernameResponse{}, //count = 0 implied
		},
	}

	f, open_err := os.Open("/go/src/owly-server/token.txt")
	check(open_err, "Could not open token file")

	accessToken, readall_err := ioutil.ReadAll(f)
	check(readall_err, "Error while reading token file")

	md := metadata.Pairs("authorization", string(accessToken))
	ctx := metadata.NewIncomingContext(context.Background(), md)

	for _, tt := range tests {
		resp, err := s.SearchUserByUsername(ctx, &tt.req)

		if err != nil {
			t.Errorf("SearchUserByUsername got unexpected error %v", err)
		} else if resp != nil && (reflect.TypeOf(resp) != reflect.TypeOf(&tt.want) ||
			resp.Count != tt.want.Count) {
			t.Errorf(
				"SearchUserByUsername(%v)=\n%v \nwanted %v",
				&tt.req, resp, &tt.want,
			)
		} else {
			log.Println("[Successfully passed TestSearchUserByUsername]")
		}
	}
}

func TestGetUserInfos(t *testing.T) {
	s := server{}

	f, openErr := os.Open("/go/src/owly-server/token.txt")
	check(openErr, "Could not open token file")

	accessToken, readallErr := ioutil.ReadAll(f)
	check(readallErr, "Error while reading token file")

	md := metadata.Pairs("authorization", string(accessToken))
	ctx := metadata.NewIncomingContext(context.Background(), md)

	searchUserResp, searchUserErr := s.SearchUserByUsername(
		ctx,
		&userpb.SearchUserByUsernameRequest{Username: "user"},
	)
	
	if searchUserErr != nil {
		t.Errorf("SearchUserByUsername got unexpected error %v", searchUserErr)
	}

	userID := searchUserResp.Results[0].Uuid

	tests := []struct {
		req  userpb.GetUserInfosRequest
		want interface{}
	}{
		{
			req: userpb.GetUserInfosRequest{
				Id: userID,
			},
			want: userpb.GetUserInfosResponse{
				Username: "user",
				Firstname: "Sample",
				Lastname: "User",
				Email: "sample-user@example",
			},
		},
		{
			req: userpb.GetUserInfosRequest{
				Id: "id that does not exist",
			},
			want: status.Error(codes.NotFound, "Error: 404 Not Found: User not found"),
		},
	}

	for _, tt := range tests {
		resp, err := s.GetUserInfos(ctx, &tt.req)

		if reflect.TypeOf(tt.want) == reflect.TypeOf(err){
			assert.Equal(t, tt.want, err)
		} else if err != nil {
			t.Errorf("GetUserInfos got unexpected error %v", err)
		} else {
			assert.Equal(t, tt.want.(userpb.GetUserInfosResponse).Username, resp.Username)
			assert.Equal(t, tt.want.(userpb.GetUserInfosResponse).Firstname, resp.Firstname)
			assert.Equal(t, tt.want.(userpb.GetUserInfosResponse).Lastname, resp.Lastname)
			assert.Equal(t, tt.want.(userpb.GetUserInfosResponse).Email, resp.Email)
		}
	}
}