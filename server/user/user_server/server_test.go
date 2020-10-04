package user_server

import (
	"reflect"
	"io/ioutil"
	"os"
	"context"
	"log"
	"primitivofr/owly/server/user/userpb"
	"testing"

	"google.golang.org/grpc/metadata"
)

// "context"
// "log"
// "os"
// "primitivofr/owly/user/auth/userpb"
// "reflect"
// "testing"

// "google.golang.org/grpc/codes"
// "google.golang.org/grpc/status"

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
			req: userpb.SearchUserByUsernameRequest{
				Username: "notexist",
			},
			want: userpb.SearchUserByUsernameResponse{
				Count: 0,
			},
		},
	}

	f, _ := os.Open("../../token.txt")
	accessToken, _ := ioutil.ReadAll(f)
	
	md := metadata.Pairs("authorization", string(accessToken))
	ctx := metadata.NewIncomingContext(context.Background(), md)

	for _, tt := range tests {
		resp, err := s.SearchUserByUsername(ctx, &tt.req)

		if err != nil {
			t.Errorf("SearchUserByUsername got unexpected error %v", err)
		} else if resp != nil && (
			reflect.TypeOf(resp) != reflect.TypeOf(&tt.want) ||
			resp.Count != tt.want.Count) {
			t.Errorf(
				"SearchUserByUsername(%v)=%v, wanted %v,",
				&tt.req, resp, &tt.want,
			)
		} else {
			log.Println("[Successfully passed TestSearchUserByUsername]")
		}
	}
}
