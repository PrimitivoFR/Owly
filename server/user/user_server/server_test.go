package user_server

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"primitivofr/owly/server/user/userpb"
	"reflect"
	"testing"

	"google.golang.org/grpc/metadata"
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
