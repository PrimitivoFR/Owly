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

	conn.Close()

}

func prepareLeaveAndDeleteChatroomTestCtx() {
	// Init client
	conn, err := grpc.Dial("server:50054", grpc.WithInsecure())
	check(err, "Error while trying to dial auth MS")
	c := authpb.NewAuthServiceClient(conn)

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

	if err := conn.Close(); err != nil {
		panic(err)
	}

	if err := conn2.Close(); err != nil {
		panic(err)
	}

}

func assertOld(t *testing.T, expected interface{}, test interface{}) {
	if reflect.DeepEqual(expected, test) {
		t.Errorf(
			"Assertion failed:\n expected\t %v of (%v)\n got\t\t %v (%v)",
			expected, reflect.TypeOf(expected), test, reflect.TypeOf(test),
		)
	}
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
		resp, err := s.CreateChatroom(appliNHCtxInc, &tt.req)
		assert.Nil(t, err, "CreateChatroom got unexpected error")

		assert.Equal(t, tt.want.Success, resp.Success, "CreateChatroom got unexpected response")
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
				Count:   3,
			},
		},
	}

	for _, tt := range tests {
		resp, err := s.GetChatroomsByUser(appliNHCtxInc, &tt.req)

		assert.Nil(t, err, "GetChatroomsByUser got unexpected error")
		assert.Equal(t, tt.want.Success, resp.Success, "GetChatroomsByUser got unexpected response")
		assert.Equal(t, tt.want.Count, resp.Count, "GetChatroomsByUser got unexpected response")
	}
}

func TestTransferOwnership(t *testing.T) {
	// Basically creates a chatroom by Toto with applinh inside, and also the opposite
	// and fills the chatroomIdByAppliNH and chatroomIdByToto
	// so this is useful here, I felt like reusing it :)
	// ... omg we need to upgrade our testing model, cause it makes us do so much shit.. Can't wait for rd-testing-model to be done, approved, merged and applied everywhere
	createUserToto()
	prepareLeaveAndDeleteChatroomTestCtx()

	s := server{}

	ctx := prepareTestingCtx()

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
			want: status.Error(codes.PermissionDenied, ""),
		},
	}

	for _, tt := range tests {
		resp, err := s.TransferOwnership(ctx, &tt.req)
		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error

			assert(t, tt.want, err)

		} else if err != nil {
			check(err, "TransferOwnership got unexpected error")
		} else {
			assert(t, tt.want, &resp)
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
				ChatroomId: chatroomIdByAppliNH,
			},
			want: status.Error(codes.PermissionDenied, ""),
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
		resp, err := s.LeaveChatroom(appliNHCtxInc, &tt.req)
		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error

			assertOld(t, tt.want, err)

		} else if err != nil {
			common_testing.CheckErr(err, "LeaveChatroom got unexpected error")
		} else {
			assertOld(t, tt.want, &resp)
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
			want: status.Error(codes.PermissionDenied, ""),
		},
		{
			req: chatroompb.DeleteChatroomRequest{
				ChatroomId: chatroomIdByAppliNH,
			},
			want: chatroompb.DeleteChatroomResponse{
				Success: true,
			},
		},
	}

	for _, tt := range tests {
		resp, err := s.DeleteChatroom(appliNHCtxInc, &tt.req)
		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error

			assertOld(t, tt.want, err)

		} else if err != nil {
			common_testing.CheckErr(err, "DeleteChatroom got unexpected error")
		} else {
			assertOld(t, tt.want, &resp)
		}

	}

}
