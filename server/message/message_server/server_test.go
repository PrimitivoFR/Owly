package message_server

import (
	"context"
	"log"
	"sort"

	"reflect"
	"testing"

	"github.com/primitivofr/owly/server/chatroom/chatroompb"
	common_jwt "github.com/primitivofr/owly/server/common/jwt"
	common_testing "github.com/primitivofr/owly/server/common/testing"
	"github.com/primitivofr/owly/server/message/messagepb"

	"time"

	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

// Contexts
var appliNHCtxInc context.Context
var totoCtxInc context.Context

var appliNHCtxOut context.Context
var totoCtxOut context.Context

const bufSize = 1024 * 1024

var lis *bufconn.Listener

var currentContext context.Context

var currentMessageId string
var appliNHMessageId string
var totoMessageId string

var uuidAppliNH string
var uuidToto string

var totoChatroomId string

func check(e error, desc string) {
	if e != nil {
		log.Println(desc)
		panic(e)
	}
}

func init() {
	// This function will automatically run first.

	// Starting chatroom MS
	// go chatroom_server.StartServer()

	// Starting message MS
	// go StartServer()

	// Starting auth MS
	// go authserver.StartServer()
	time.Sleep(2 * time.Second)

	var err error
	appliNHCtxInc, appliNHCtxOut = common_testing.PrepareCtxWithToken("applinh")
	totoCtxInc, totoCtxOut = common_testing.PrepareCtxWithToken("toto")

	uuidAppliNH, err = common_jwt.ReadUUIDFromContext(appliNHCtxInc)
	common_testing.CheckErr(err, "Error reading uuid for applinh")

	uuidToto, err = common_jwt.ReadUUIDFromContext(totoCtxInc)
	common_testing.CheckErr(err, "Error reading uuid for toto")

	// Create chatroom by toto
	reqCreateChatroom := &chatroompb.CreateChatroomRequest{
		Name:  "ChatroomYEH",
		Users: []string{uuidAppliNH},
	}

	resCreateChatroom := &chatroompb.CreateChatroomResponse{}

	_, err = common_testing.ContextGenerator(
		totoCtxOut,
		"chatroom.ChatroomService",
		"CreateChatroom",
		"50052",
		reqCreateChatroom,
		resCreateChatroom)

	common_testing.CheckErr(err, "Error dialing with CreateChatroom")

	totoChatroomId = resCreateChatroom.ID

	// Send a message by toto
	reqSendMessage := &messagepb.SendMessageRequest{
		Message: &messagepb.Message{
			Content:    "boi",
			ChatroomID: totoChatroomId,
			AuthorNAME: "toto",
		},
	}

	resSendMessage := &messagepb.SendMessageResponse{}

	_, err = common_testing.ContextGenerator(
		totoCtxOut,
		"message.MessageService",
		"SendMessage",
		"50053",
		reqSendMessage,
		resSendMessage)

	common_testing.CheckErr(err, "Error dialing with SendMessage")
	totoMessageId = resSendMessage.Id

}

func TestSendMessage(t *testing.T) {

	s := server{}

	tests := []struct {
		req  messagepb.SendMessageRequest
		want messagepb.SendMessageResponse
	}{
		{
			req: messagepb.SendMessageRequest{
				Message: &messagepb.Message{
					ChatroomID: totoChatroomId,
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
		res, err := s.SendMessage(appliNHCtxInc, &tt.req)
		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error

			common_testing.CmpAssertEqual(t, tt.want, err, cmpopts.EquateErrors())

		} else {
			if o := assert.Nil(t, err, "SendMessage got unexpected error"); o {
				common_testing.CmpAssertEqual(t, tt.want, *res, cmpopts.IgnoreUnexported(*res), cmpopts.IgnoreFields(*res, "Id"))
			}
		}

	}

}

func TestGetMessagesByChatroom(t *testing.T) {
	s := server{}

	tests := []struct {
		req  messagepb.GetMessagesByChatroomRequest
		want messagepb.GetMessagesByChatroomResponse
	}{
		{
			req: messagepb.GetMessagesByChatroomRequest{
				ChatroomID: totoChatroomId,
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
		res, err := s.GetMessagesByChatroom(appliNHCtxInc, &tt.req)

		// Sorting messages list from old to new
		sort.SliceStable(res.Messages, func(i, j int) bool {
			return res.Messages[i].Timestamp < res.Messages[j].Timestamp
		})

		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error

			common_testing.CmpAssertEqual(t, tt.want, err, cmpopts.EquateErrors())

		} else {
			if o := assert.Nil(t, err, "GetMessagesByChatroom got unexpected error"); o {
				common_testing.CmpAssertEqual(t, tt.want.Messages[0].Content, res.Messages[1].Content)
				common_testing.CmpAssertEqual(t, tt.want.Messages[0].AuthorNAME, res.Messages[1].AuthorNAME)
			}
		}
		appliNHMessageId = res.Messages[1].Id
	}
}

func TestUpdateMessageContent(t *testing.T) {

	s := server{}

	tests := []struct {
		req  messagepb.UpdateMessageContentRequest
		want interface{}
	}{
		{
			req: messagepb.UpdateMessageContentRequest{
				MessageId:  appliNHMessageId,
				ChatroomId: totoChatroomId,
				NewContent: "Slt test",
			},
			want: messagepb.UpdateMessageContentResponse{
				Success: true,
				Message: &messagepb.Message{
					Id: appliNHMessageId,
				},
			},
		},
		{
			req: messagepb.UpdateMessageContentRequest{
				ChatroomId: totoChatroomId,
				MessageId:  totoMessageId,
				NewContent: "ah ben nn enfait",
			},
			want: status.Error(codes.PermissionDenied, "This user "+uuidAppliNH+" is not the author of the message "+totoMessageId+". He can't do anything with it"),
		},
	}

	for _, tt := range tests {
		res, err := s.UpdateMessageContent(appliNHCtxInc, &tt.req)
		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error

			common_testing.CmpAssertEqual(t, tt.want, err, cmpopts.EquateErrors())

		} else {

			if o := assert.Nil(t, err, "UpdateMessageContent got unexpected error"); o {
				common_testing.CmpAssertEqual(t, tt.want, *res, cmpopts.IgnoreUnexported(*res, *res.Message))
			}

		}

	}
}

func TestDeleteMessage(t *testing.T) {

	s := server{}

	tests := []struct {
		req  messagepb.DeleteMessageRequest
		want messagepb.DeleteMessageResponse
	}{
		{
			req: messagepb.DeleteMessageRequest{
				ChatroomID: totoChatroomId,
				MessageID:  appliNHMessageId,
			},
			want: messagepb.DeleteMessageResponse{
				Success: true,
			},
		},
	}

	for _, tt := range tests {
		res, err := s.DeleteMessage(appliNHCtxInc, &tt.req)
		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error

			common_testing.CmpAssertEqual(t, tt.want, err, cmpopts.EquateErrors())
		} else {

			if o := assert.Nil(t, err, "DeleteMessage got unexpected error"); o {
				common_testing.CmpAssertEqual(t, tt.want, *res, cmpopts.IgnoreUnexported(*res))
			}
		}
	}
}
