package chatroom_server

import (
	"context"
	"encoding/json"
	authserver "primitivofr/owly/server/auth/auth_server"
	"primitivofr/owly/server/auth/authpb"
	"primitivofr/owly/server/chatroom/chatroompb"
	common_jwt "primitivofr/owly/server/common/jwt"
	common_testing "primitivofr/owly/server/common/testing"

	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Contexts
var appliNHCtxInc context.Context
var totoCtxInc context.Context

var appliNHCtxOut context.Context
var totoCtxOut context.Context

// Misc
var chatroomIdByToto string
var chatroomIdByAppliNH string
var chatroomIdByAppliNHBIS string

var uuidAppliNH string
var uuidToto string

func init() {

	appliNHCtxInc, appliNHCtxOut = common_testing.PrepareCtxWithToken("applinh")

	//---- CREATE ANOTHER USER
	// Starting auth MS
	go authserver.StartServer()

	time.Sleep(1 * time.Second)

	reqCreateUser := &authpb.CreateNewUserRequest{
		Email:     "toto",
		FirstName: "toto",
		LastName:  "toto",
		Username:  "toto",
		Password:  "toto",
	}

	resCreateUser := &authpb.CreateNewUserResponse{}

	_, err := common_testing.ContextGenerator(
		context.Background(),
		"auth.AuthService",
		"CreateNewUser",
		"50054",
		reqCreateUser,
		resCreateUser)

	common_testing.CheckErr(err, "")

	b, err := json.Marshal(reqCreateUser)
	common_testing.CheckErr(err, "")

	common_testing.SaveToKvdb("users", reqCreateUser.Username, string(b))

	// Login
	reqLogin := &authpb.LoginUserRequest{
		Username: "toto",
		Password: "toto",
	}
	resLogin := &authpb.LoginUserResponse{}

	_, err = common_testing.ContextGenerator(
		context.Background(),
		"auth.AuthService",
		"LoginUser",
		"50054",
		reqLogin,
		resLogin)

	common_testing.CheckErr(err, "")

	common_testing.SaveToKvdb("tokens", reqLogin.Username, resLogin.Result.AccessToken)

	totoCtxInc, totoCtxOut = common_testing.PrepareCtxWithToken(reqLogin.Username)

	//---- CREATE NEW CHATROOMS

	// Starting chatroom MS
	go StartServer()
	time.Sleep(1 * time.Second)

	// Retrieving AppliNH uuid
	uuidAppliNH, err = common_jwt.ReadUUIDFromContext(appliNHCtxInc)
	common_testing.CheckErr(err, "Err while reading uuid from applinh token")

	// Create Chatroom with Toto, and AppliNH inside
	reqCreateChatroom := &chatroompb.CreateChatroomRequest{
		Name:  "ChatroomByToto",
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

	common_testing.CheckErr(err, "")

	chatroomIdByToto = resCreateChatroom.ID

	// Retrieving toto uuid
	uuidToto, err = common_jwt.ReadUUIDFromContext(totoCtxInc)
	common_testing.CheckErr(err, "Err while reading uuid from toto token")

	// Create Chatroom with AppliNH, and toto inside
	reqCreateChatroom2 := &chatroompb.CreateChatroomRequest{
		Name:  "ChatroomByAppliNH",
		Users: []string{uuidToto},
	}
	resCreateChatroom2 := &chatroompb.CreateChatroomResponse{}

	_, err = common_testing.ContextGenerator(
		appliNHCtxOut,
		"chatroom.ChatroomService",
		"CreateChatroom",
		"50052",
		reqCreateChatroom2,
		resCreateChatroom2)

	common_testing.CheckErr(err, "")

	chatroomIdByAppliNH = resCreateChatroom2.ID

	// Create second Chatroom with AppliNH, and toto inside
	reqCreateChatroom3 := &chatroompb.CreateChatroomRequest{
		Name:  "ChatroomByAppliNH_BIS",
		Users: []string{uuidToto},
	}
	resCreateChatroom3 := &chatroompb.CreateChatroomResponse{}

	_, err = common_testing.ContextGenerator(
		appliNHCtxOut,
		"chatroom.ChatroomService",
		"CreateChatroom",
		"50052",
		reqCreateChatroom3,
		resCreateChatroom3)

	common_testing.CheckErr(err, "")

	chatroomIdByAppliNHBIS = resCreateChatroom3.ID

}

func TestCreateChatroom(t *testing.T) {
	s := server{}

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
		res, err := s.CreateChatroom(appliNHCtxInc, &tt.req)
		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error
			common_testing.CmpAssertEqual(t, tt.want, err, cmpopts.EquateErrors())

		} else {
			if o := assert.Nil(t, err, "CreateChatroom got unexpected error"); o {
				common_testing.CmpAssertEqual(t, tt.want, *res, cmpopts.IgnoreUnexported(*res), cmpopts.IgnoreFields(*res, "ID"))
			}
		}

	}
}

func TestGetChatroomsByUser(t *testing.T) {
	s := server{}

	tests := []struct {
		req  chatroompb.GetChatroomsByUserRequest
		want chatroompb.GetChatroomsByUserResponse
	}{
		{
			req: chatroompb.GetChatroomsByUserRequest{},
			want: chatroompb.GetChatroomsByUserResponse{
				Success: true,
				Count:   4,
			},
		},
	}

	for _, tt := range tests {
		res, err := s.GetChatroomsByUser(appliNHCtxInc, &tt.req)

		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error
			common_testing.CmpAssertEqual(t, tt.want, err, cmpopts.EquateErrors())
		} else {
			if o := assert.Nil(t, err, "GetChatroomsByUser got unexpected error"); o {
				common_testing.CmpAssertEqual(t, tt.want, *res, cmpopts.IgnoreUnexported(*res), cmpopts.IgnoreFields(*res, "Chatrooms"))
			}
		}

	}
}

func TestTransferOwnership(t *testing.T) {

	s := server{}

	tests := []struct {
		req  chatroompb.TranferOwnershipRequest
		want interface{}
	}{
		{
			req: chatroompb.TranferOwnershipRequest{
				ChatroomId: chatroomIdByAppliNH,
				NewOwnerId: uuidToto,
			},
			want: chatroompb.TranferOwnershipResponse{
				Success: true,
			},
		},

		{
			req: chatroompb.TranferOwnershipRequest{
				ChatroomId: chatroomIdByToto,
				NewOwnerId: uuidAppliNH,
			},
			want: status.Error(codes.PermissionDenied, "Current user is not the rightful owner of this chatroom, and thus cannot transfer its ownership."),
		},
	}

	for _, tt := range tests {
		res, err := s.TransferOwnership(appliNHCtxInc, &tt.req)
		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error
			common_testing.CmpAssertEqual(t, tt.want, err, cmpopts.EquateErrors())
		} else {
			if o := assert.Nil(t, err, "TransferOwnership got unexpected error"); o {
				common_testing.CmpAssertEqual(t, tt.want, *res, cmpopts.IgnoreUnexported(*res))
			}
		}

	}

}

func TestLeaveChatroom(t *testing.T) {

	s := server{}

	tests := []struct {
		req  chatroompb.LeaveChatroomRequest
		want interface{}
	}{
		{
			req: chatroompb.LeaveChatroomRequest{
				ChatroomId: chatroomIdByAppliNHBIS,
			},
			want: status.Error(codes.PermissionDenied, "Current user is the rightful owner of this chatroom, and thus cannot leave it."),
		},

		{
			req: chatroompb.LeaveChatroomRequest{
				ChatroomId: chatroomIdByToto,
			},
			want: chatroompb.LeaveChatroomResponse{
				Success: true,
			},
		},
	}

	for _, tt := range tests {
		res, err := s.LeaveChatroom(appliNHCtxInc, &tt.req)
		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error
			common_testing.CmpAssertEqual(t, tt.want, err, cmpopts.EquateErrors())

		} else {
			if o := assert.Nil(t, err, "LeaveChatroom got unexpected error"); o {
				common_testing.CmpAssertEqual(t, tt.want, *res, cmpopts.IgnoreUnexported(*res))
			}
		}

	}

}
func TestDeleteChatroom(t *testing.T) {

	s := server{}

	tests := []struct {
		req  chatroompb.DeleteChatroomRequest
		want interface{}
	}{
		{
			req: chatroompb.DeleteChatroomRequest{
				ChatroomId: chatroomIdByToto,
			},
			want: status.Error(codes.PermissionDenied, "Current user is not the rightful owner of this chatroom, and thus cannot delete it."),
		},
		{
			req: chatroompb.DeleteChatroomRequest{
				ChatroomId: chatroomIdByAppliNHBIS,
			},
			want: chatroompb.DeleteChatroomResponse{
				Success: true,
			},
		},
	}

	for _, tt := range tests {

		res, err := s.DeleteChatroom(appliNHCtxInc, &tt.req)

		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error
			common_testing.CmpAssertEqual(t, tt.want, err, cmpopts.EquateErrors())

		} else {
			if o := assert.Nil(t, err, "DeleteChatroom got unexpected error"); o {
				common_testing.CmpAssertEqual(t, tt.want, *res, cmpopts.IgnoreUnexported(*res))
			}
		}

	}

}
