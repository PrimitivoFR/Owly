package message_server

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	authserver "primitivofr/owly/server/auth/auth_server"
	"primitivofr/owly/server/auth/authpb"
	"sort"

	"primitivofr/owly/server/chatroom/chatroom_server"
	"primitivofr/owly/server/chatroom/chatroompb"
	common_jwt "primitivofr/owly/server/common/jwt"
	"primitivofr/owly/server/message/messagepb"
	"reflect"
	"testing"

	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

var currentContext context.Context

var currentChatroomId string
var currentMessageId string
var totoMessageId string

var uuidAppliNH string

var cMessage messagepb.MessageServiceClient

func assert(t *testing.T, expected interface{}, test interface{}) {
	fmt.Println(expected, test)
	// v is the interface{}
	// v := reflect.ValueOf(&test).Elem()

	// // Allocate a temporary variable with type of the struct.
	// //    v.Elem() is the vale contained in the interface.
	// tmp := reflect.New(v.Elem().Type()).Elem()

	// // Copy the struct value contained in interface to
	// // the temporary variable.
	// tmp.Set(v.Elem())

	// // Set the field.
	// state := tmp.FieldByName("state")
	// ptrToY := unsafe.Pointer(state.UnsafeAddr())

	// realPtrToY := (*interface{})(ptrToY)
	// *realPtrToY = nil
	// // Set the interface to the modified struct value.
	// v.Set(tmp)

	if !(reflect.DeepEqual(expected, test)) {
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

func init() {
	// This function will automatically run first.

	// Starting chatroom MS
	go chatroom_server.StartServer()

	// Starting message MS
	go StartServer()

	// Starting auth MS
	go authserver.StartServer()
	time.Sleep(2 * time.Second)

	time.Sleep(2 * time.Second)
	prepareTestingCtx()
	time.Sleep(1 * time.Second)

	// Connections
	connAuth, err := grpc.Dial("server:50054", grpc.WithInsecure())
	check(err, "Error whil trying to dial auth MS")

	connMessage, err := grpc.Dial("server:50053", grpc.WithInsecure())
	check(err, "Error whil trying to dial message MS")

	connChatroom, err := grpc.Dial("server:50052", grpc.WithInsecure())
	check(err, "Error whil trying to dial chatroom MS")

	// Contexts and token
	f, openErr := os.Open("/go/src/owly-server/token.txt")
	check(openErr, "Could not open token file")
	accessToken, err := ioutil.ReadAll(f)
	check(err, "Error while reading token file")
	uuidAppliNH, err = common_jwt.ExtractUUIDfromJWT(string(accessToken))
	check(err, "Err while reading uuid from applinh token")

	// Login toto
	reqLogin := &authpb.LoginUserRequest{
		Username: "toto",
		Password: "toto",
	}

	cAuth := authpb.NewAuthServiceClient(connAuth)
	resLoginToto, err := cAuth.LoginUser(context.Background(), reqLogin)
	check(err, "Error while logging in with user toto")

	mdToto := metadata.Pairs("authorization", string(resLoginToto.Result.AccessToken))
	ctxToto := metadata.NewOutgoingContext(context.Background(), mdToto)

	// Create chatroom by toto
	reqChatroom := &chatroompb.CreateChatroomRequest{
		Name:  "ChatroomYEH",
		Users: []string{uuidAppliNH},
	}
	cChatroom := chatroompb.NewChatroomServiceClient(connChatroom)
	res, err := cChatroom.CreateChatroom(ctxToto, reqChatroom)
	currentChatroomId = res.ID

	// Send a message by toto
	reqMessage := &messagepb.SendMessageRequest{
		Message: &messagepb.Message{
			Content:    "boi",
			ChatroomID: currentChatroomId,
			AuthorNAME: "toto",
		},
	}
	cMessage = messagepb.NewMessageServiceClient(connMessage)
	_, err = cMessage.SendMessage(ctxToto, reqMessage)
	check(err, "Error while sending message with user toto")

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
		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error

			assert(t, tt.want, err)

		} else {
			if o := cmp.Equal(tt.want, *res, cmpopts.IgnoreUnexported(*res)); o == false {
				t.Errorf(
					"Assertion failed:\n expected\t %v of (%v)\n got\t\t %v (%v)",
					tt.want, reflect.TypeOf(tt.want), res, reflect.TypeOf(res),
				)
			}

			//assert(t, tt.want, res)
		}

	}
	time.Sleep(2 * time.Second)

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

		fmt.Println(res.Messages[0])
		fmt.Println(res.Messages[1])

		// Sorting messages list from old to new
		sort.SliceStable(res.Messages, func(i, j int) bool {
			return res.Messages[i].Timestamp < res.Messages[j].Timestamp
		})

		check(err, "Error while trying to get messages by chatroom")
		assert(t, tt.want.Messages[0].Content, res.Messages[1].Content)
		assert(t, tt.want.Messages[0].AuthorNAME, res.Messages[1].AuthorNAME)

		currentMessageId = res.Messages[1].Id
		totoMessageId = res.Messages[0].Id
	}
}

func TestUpdateMessageContent(t *testing.T) {
	tests := []struct {
		req  messagepb.UpdateMessageContentRequest
		want interface{}
	}{
		{
			req: messagepb.UpdateMessageContentRequest{
				MessageId:  currentMessageId,
				ChatroomId: currentChatroomId,
				NewContent: "Slt test",
			},
			want: &messagepb.UpdateMessageContentResponse{
				Success: true,
				Message: &messagepb.Message{
					Id: currentMessageId,
				},
			},
		},
		{
			req: messagepb.UpdateMessageContentRequest{
				ChatroomId: currentChatroomId,
				MessageId:  totoMessageId,
				NewContent: "ah ben nn enfait",
			},
			want: status.Error(codes.PermissionDenied, "This user "+uuidAppliNH+" is not the author of the message "+totoMessageId+". He can't do anything with it"),
		},
	}

	for _, tt := range tests {
		res, err := cMessage.UpdateMessageContent(currentContext, &tt.req)
		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error

			assert(t, tt.want, err)

		} else {

			assert(t, tt.want, res)
		}

	}
}

func TestDeleteMessage(t *testing.T) {

	tests := []struct {
		req  messagepb.DeleteMessageRequest
		want *messagepb.DeleteMessageResponse
	}{
		{
			req: messagepb.DeleteMessageRequest{
				ChatroomID: currentChatroomId,
				MessageID:  currentMessageId,
			},
			want: &messagepb.DeleteMessageResponse{
				Success: true,
			},
		},
	}

	for _, tt := range tests {
		res, err := cMessage.DeleteMessage(currentContext, &tt.req)
		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error

			assert(t, tt.want, err)

		} else {
			assert(t, tt.want, res)
		}
	}
}
