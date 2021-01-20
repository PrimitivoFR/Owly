package user_server

import (
	"context"
	// "io/ioutil"
	// "log"
	// "os"
	"github.com/primitivofr/owly/server/user/userpb"
	"reflect"
	"testing"

	//common_jwt "github.com/primitivofr/owly/server/common/jwt"
	common_testing "github.com/primitivofr/owly/server/common/testing"

	"google.golang.org/grpc/codes"
	// "google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

// Contexts
var appliNHCtxInc context.Context

// Misc
func init() {
	appliNHCtxInc, _ = common_testing.PrepareCtxWithToken("applinh")
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

	for _, tt := range tests {
		res, err := s.SearchUserByUsername(appliNHCtxInc, &tt.req)

		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error
			common_testing.CmpAssertEqual(t, tt.want, err, cmpopts.EquateErrors())
		} else {
			if o := assert.Nil(t, err, "SearchUserByUsername got unexpected error"); o {
				assert.Equal(t, tt.want.Count, res.Count)
			}
		}
	}
}

func TestGetUserInfos(t *testing.T) {
	s := server{}

	searchUserResp, searchUserErr := s.SearchUserByUsername(
		appliNHCtxInc,
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
				Username:  "user",
				Firstname: "Sample",
				Lastname:  "User",
				Email:     "sample-user@example",
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
		res, err := s.GetUserInfos(appliNHCtxInc, &tt.req)

		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error
			common_testing.CmpAssertEqual(t, tt.want, err, cmpopts.EquateErrors())
		} else {
			if o := assert.Nil(t, err, "SearchUserByUsername got unexpected error"); o {
				common_testing.CmpAssertEqual(t, tt.want, *res, cmpopts.IgnoreUnexported(*res))
			}
		}
	}
}
