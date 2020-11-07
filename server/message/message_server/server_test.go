package message_server

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"primitivofr/owly/server/chatroom/chatroom_server"
	"primitivofr/owly/server/chatroom/chatroompb"
	"primitivofr/owly/server/message/messagepb"
	"reflect"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

var currentContext context.Context

var currentChatroomId string
var currentMessageId string

var cMessage messagepb.MessageServiceClient

func assert(t *testing.T, expected interface{}, test interface{}) {
	if test != expected {
		t.Errorf(
			"Assertion failed:\n expected\t %v of (%v)\n got\t\t %v (%v)",
			expected, reflect.TypeOf(expected), test, reflect.TypeOf(test),
		)
	}
}

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
	log.Println("accessToken: ", string(accessToken))
	md := metadata.Pairs("authorization", string(accessToken))
	currentContext = metadata.NewOutgoingContext(context.Background(), md)
	return
}

func getChatroomId() {
	conn, err := grpc.Dial("server:50052", grpc.WithInsecure())
	check(err, "Error whil trying to dial chatroom MS")
	c := chatroompb.NewChatroomServiceClient(conn)

	res, err := c.GetChatroomsByUser(currentContext, &chatroompb.GetChatroomsByUserRequest{})
	check(err, "Error while trying to get chatrooms")

	currentChatroomId = res.Chatrooms[0].Id

}

func init() {
	// This function will automatically run first.

	go chatroom_server.StartServer()
	go StartServer()
	time.Sleep(2 * time.Second)
	prepareTestingCtx()
	time.Sleep(1 * time.Second)
	getChatroomId()

	conn, err := grpc.Dial("server:50053", grpc.WithInsecure())
	check(err, "Error whil trying to dial message MS")
	cMessage = messagepb.NewMessageServiceClient(conn)

}

func TestSendMessage(t *testing.T) {

	tests := []struct {
		req  messagepb.SendMessageRequest
		want messagepb.SendMessageResponse
	}{
		{
			req: messagepb.SendMessageRequest{
				Message: &messagepb.Message{
					ChatroomID: currentChatroomId,
					Content:    "test",
					AuthorNAME: "AppliNH",
				},
			},
			want: messagepb.SendMessageResponse{
				Success: true,
			},
		},
	}

	for _, tt := range tests {
		res, err := cMessage.SendMessage(currentContext, &tt.req)
		check(err, "Error while trying to send a message")
		assert(t, tt.want.Success, res.Success)

	}

}

func TestGetMessagesByChatroom(t *testing.T) {

	tests := []struct {
		req  messagepb.GetMessagesByChatroomRequest
		want messagepb.GetMessagesByChatroomResponse
	}{
		{
			req: messagepb.GetMessagesByChatroomRequest{
				ChatroomID: currentChatroomId,
			},
			want: messagepb.GetMessagesByChatroomResponse{
				Messages: []*messagepb.Message{
					&messagepb.Message{
						Content:    "test",
						AuthorNAME: "AppliNH",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		res, err := cMessage.GetMessagesByChatroom(currentContext, &tt.req)
		check(err, "Error while trying to get messages by chatroom")
		assert(t, tt.want.Messages[0].Content, res.Messages[0].Content)
		assert(t, tt.want.Messages[0].AuthorNAME, res.Messages[0].AuthorNAME)

		currentMessageId = res.Messages[0].Id
	}
}

func TestUpdateMessageContent(t *testing.T) {
	tests := []struct {
		req		messagepb.UpdateMessageContentRequest
		want	messagepb.UpdateMessageContentResponse
	}{
		{
			req: messagepb.UpdateMessageContentRequest{
			MessageId: currentMessageId,
			ChatroomId: currentChatroomId,
			NewContent: "Slt test",
		},
		want: messagepb.UpdateMessageContentResponse{
			Success: true,
			Message: &messagepb.Message{
				Content: "Slt test",
				AuthorNAME: "AppliNH",
				},
			},
		},
	}

	for _, tt := range tests {
		res, err := cMessage.UpdateMessageContent(currentContext, &tt.req)
		check(err, "Error while trying to update message")
		assert(t, tt.want.Success, res.Success)
		assert(t, tt.want.Message.Content, res.Message.Content)
		assert(t, tt.want.Message.AuthorNAME, res.Message.AuthorNAME)
	}
}


func TestDeleteMessage(t *testing.T) {

	tests := []struct {
		req  messagepb.DeleteMessageRequest
		want messagepb.DeleteMessageResponse
	}{
		{
			req: messagepb.DeleteMessageRequest{
				ChatroomID: currentChatroomId,
				MessageID:  currentMessageId,
			},
			want: messagepb.DeleteMessageResponse{
				Success: true,
			},
		},
	}

	for _, tt := range tests {
		res, err := cMessage.DeleteMessage(currentContext, &tt.req)
		check(err, "Error while trying to delete a message")
		assert(t, tt.want.Success, res.Success)
	}
}
