package chatroom_server

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	authserver "primitivofr/owly/server/auth/auth_server"
	"primitivofr/owly/server/auth/authpb"
	"primitivofr/owly/server/chatroom/chatroompb"
	common_jwt "primitivofr/owly/server/common/jwt"
	common_mongo "primitivofr/owly/server/common/mongo"
	"reflect"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
	if reflect.DeepEqual(expected, test) {
		t.Errorf(
			"Assertion failed:\n expected\t %v of (%v)\n got\t\t %v (%v)",
			expected, reflect.TypeOf(expected), test, reflect.TypeOf(test),
		)
	}
}

var chatroomIdByToto string
var chatroomIdByAppliNH string

func prepareLeaveAndDeleteChatroomTestCtx() {
	log.Println("prepareLeaveAndDeleteChatroomTestCtx")
	//---- CREATE ANOTHER USER

	// Starting auth MS
	go authserver.StartServer()
	time.Sleep(2 * time.Second)

	// Init client
	conn, err := grpc.Dial("server:50054", grpc.WithInsecure())
	check(err, "Error while trying to dial auth MS")
	c := authpb.NewAuthServiceClient(conn)

	reqCreateUser := &authpb.CreateNewUserRequest{
		Email:     "toto",
		FirstName: "toto",
		LastName:  "toto",
		Username:  "toto",
		Password:  "toto",
	}

	_, err = c.CreateNewUser(context.Background(), reqCreateUser)
	check(err, "Error while creating user toto")

	// Login
	reqLogin := &authpb.LoginUserRequest{
		Username: "toto",
		Password: "toto",
	}
	resLogin, err := c.LoginUser(context.Background(), reqLogin)
	check(err, "Error while logging in with user toto")

	totoAccessToken := resLogin.Result.AccessToken

	mdToto := metadata.Pairs("authorization", string(totoAccessToken))
	ctxToto := metadata.NewOutgoingContext(context.Background(), mdToto)

	f, openErr := os.Open("/go/src/owly-server/token.txt")
	check(openErr, "Could not open token file")

	accessTokenAppliNH, readAllErr := ioutil.ReadAll(f)
	check(readAllErr, "Error while reading token file")

	mdAppliNH := metadata.Pairs("authorization", string(accessTokenAppliNH))
	ctxAppliNH := metadata.NewOutgoingContext(context.Background(), mdAppliNH)

	//---- CREATE NEW CHATROOM

	// Starting chatroom MS
	go StartServer()
	time.Sleep(2 * time.Second)

	// Init client
	conn2, err := grpc.Dial("server:50052", grpc.WithInsecure())
	check(err, "Error whil trying to dial chatroom MS")
	cChatroom := chatroompb.NewChatroomServiceClient(conn2)

	// ___ TOTO ___

	// Create Chatroom with Toto
	uuidAppliNH, err := common_jwt.ExtractUUIDfromJWT(string(accessTokenAppliNH))
	check(err, "Err while reading uuid from applinh token")

	reqCreateChatroom := &chatroompb.CreateChatroomRequest{
		Name:  "ChatroomByToto",
		Users: []string{uuidAppliNH},
	}
	resCreateChatroom, err := cChatroom.CreateChatroom(ctxToto, reqCreateChatroom)
	check(err, "Err while creating chatroom for owner toto")

	chatroomIdByToto = resCreateChatroom.ID

	// ___ AppliNH ___

	// Create Chatroom with AppliNH
	uuidToto, err := common_jwt.ExtractUUIDfromJWT(string(totoAccessToken))
	check(err, "Err while reading uuid from toto token")

	reqCreateChatroom2 := &chatroompb.CreateChatroomRequest{
		Name:  "ChatroomByAppliNH",
		Users: []string{uuidToto},
	}

	resCreateChatroom2, err := cChatroom.CreateChatroom(ctxAppliNH, reqCreateChatroom2)
	check(err, "Err while creating chatroom for owner AppliNH")

	chatroomIdByAppliNH = resCreateChatroom2.ID

}

func TestCreateChatroom(t *testing.T) {
	s := server{}
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

		assert(t, tt.want.Success, &resp.Success)
	}
}

func TestGetChatroomsByUser(t *testing.T) {
	s := server{}
	ctx := prepareTestingCtx()

	tests := []struct {
		req  chatroompb.GetChatroomsByUserRequest
		want chatroompb.GetChatroomsByUserResponse
	}{
		{
			req: chatroompb.GetChatroomsByUserRequest{},
			want: chatroompb.GetChatroomsByUserResponse{
				Success: true,
				Count:   1,
			},
		},
	}

	for _, tt := range tests {
		resp, err := s.GetChatroomsByUser(ctx, &tt.req)
		check(err, "GetChatroomsByUser got unexpected error")

		assert(t, tt.want.Success, &resp.Success)
		assert(t, tt.want.Count, &resp.Count)
	}
}
func TestLeaveChatroom(t *testing.T) {
	prepareLeaveAndDeleteChatroomTestCtx()

	s := server{}

	ctx := prepareTestingCtx()

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
		resp, err := s.LeaveChatroom(ctx, &tt.req)
		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error

			// wantedErr, _ := status.FromError(tt.want.(error))
			// actualErr, _ := status.FromError(err)

			assert(t, tt.want, err)

		} else if err != nil {
			check(err, "LeaveChatroom got unexpected error")
		} else {
			assert(t, tt.want, &resp)
		}

	}

}
func TestDeleteChatroom(t *testing.T) {

	s := server{}

	ctx := prepareTestingCtx()

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
		resp, err := s.DeleteChatroom(ctx, &tt.req)
		if reflect.TypeOf(tt.want) == reflect.TypeOf(err) {
			// We're expecting an error

			// wantedErr, _ := status.FromError(tt.want.(error))
			// actualErr, _ := status.FromError(err)

			assert(t, tt.want, err)

		} else if err != nil {
			check(err, "DeleteChatroom got unexpected error")
		} else {
			assert(t, tt.want, &resp)
		}

	}

}
