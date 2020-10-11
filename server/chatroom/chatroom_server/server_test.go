package chatroom_server

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"primitivofr/owly/server/chatroom/chatroompb"
	common_mongo "primitivofr/owly/server/common/mongo"
	"testing"
	"reflect"

	"google.golang.org/grpc/metadata"
)

func check(e error, desc string) {
	if e != nil {
		log.Println(desc)
		panic(e)
	}
}

func prepareTestingCtx() (ctx context.Context) {
	f, open_err := os.Open("/go/src/owly-server/token.txt")
	check(open_err, "Could not open token file")

	accessToken, readall_err := ioutil.ReadAll(f)
	check(readall_err, "Error while reading token file")

	md := metadata.Pairs("authorization", string(accessToken))
	ctx = metadata.NewIncomingContext(context.Background(), md)
	return
}

func setubDb() {
	err := common_mongo.SetupMongoDB()
	check(err, "Error while setting up mongodb")
}

func assert(t *testing.T, expected interface{}, test interface{}) {
	if test != expected {
		t.Errorf(
			"Assertion failed:\n expected\t %v of (%v)\n got\t\t %v (%v)",
			expected, reflect.TypeOf(expected), test, reflect.TypeOf(test),
		)
	}
}

func TestCreateChatroom(t *testing.T) {
	s := server{}
	setubDb()
	ctx := prepareTestingCtx()

	tests := []struct {
		req  chatroompb.CreateChatroomRequest
		want chatroompb.CreateChatroomResponse
	}{
		{
			req: chatroompb.CreateChatroomRequest{
				Name: "testing_chatroom",
				// chatroom with only current user
			},
			want: chatroompb.CreateChatroomResponse{
				Success: true,
			},
		},
		// TODO : add test case with multiple users (but we need the user ids)
	}

	for _, tt := range tests {
		resp, err := s.CreateChatroom(ctx, &tt.req)
		check(err, "CreateChatroom got unexpected error")

		assert(t, tt.want.Success, resp.Success)
	}
}

func TestGetChatroomsByUser(t *testing.T) {
	s := server{}
	setubDb()
	ctx := prepareTestingCtx()

	tests := []struct {
		req  chatroompb.GetChatroomsByUserRequest
		want chatroompb.GetChatroomsByUserResponse
	}{
		{
			req: chatroompb.GetChatroomsByUserRequest{},
			want: chatroompb.GetChatroomsByUserResponse{
				Success: true,
				Count: 1,
			},
		},
	}

	for _, tt := range tests {
		resp, err := s.GetChatroomsByUser(ctx, &tt.req)
		check(err, "GetChatroomsByUser got unexpected error")

		assert(t, tt.want.Success, resp.Success)
		assert(t, tt.want.Count, resp.Count)
	}
}