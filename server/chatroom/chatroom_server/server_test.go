package chatroom_server

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"primitivofr/owly/server/chatroom/chatroompb"
	common_mongo "primitivofr/owly/server/common/mongo"
	"testing"

	"google.golang.org/grpc/metadata"
)

func check(e error, desc string) {
	if e != nil {
		log.Println(desc)
		panic(e)
	}
}

func prepare_testing_ctx() (ctx context.Context) {
	f, open_err := os.Open("/go/src/owly-server/token.txt")
	check(open_err, "Could not open token file")

	accessToken, readall_err := ioutil.ReadAll(f)
	check(readall_err, "Error while reading token file")

	md := metadata.Pairs("authorization", string(accessToken))
	ctx = metadata.NewIncomingContext(context.Background(), md)
	return
}

func TestCreateChatroom(t *testing.T) {
	s := server{}

	err := common_mongo.SetupMongoDB()

	check(err, "Error while setting up mongodb")

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
	}

	ctx := prepare_testing_ctx()

	for _, tt := range tests {
		resp, err := s.CreateChatroom(ctx, &tt.req)
		check(err, "CreateChatroom got unexpected error")

		if resp.Success != tt.want.Success {
			t.Errorf(
				"CreateChatroom Success state was not expected for name '%v'",
				&tt.req.Name,
			)
		}
	}
}
