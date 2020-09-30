package authserver

import (
	"context"
	"log"
	"primitivofr/owly/server/auth/authpb"
	"testing"
)

func TestCreateNewUser(t *testing.T) {

	s := server{}

	tests := []struct {
		req  authpb.CreateNewUserRequest
		want authpb.CreateNewUserResponse
	}{
		{
			req: authpb.CreateNewUserRequest{
				Email:     "toto@toto.fr",
				Password:  "Aze123",
				Username:  "applinh",
				FirstName: "Thomas",
				LastName:  "Martin",
			},
			want: authpb.CreateNewUserResponse{
				Success: true,
			},
		},
	}

	for _, tt := range tests {
		req := &tt.req

		log.Println(req)

		resp, err := s.CreateNewUser(context.Background(), req)

		if err != nil {
			t.Errorf("CreateNewUser got unexpected error %v", err)
		} else if resp.Success != tt.want.Success {
			t.Errorf("CreateNewUser(%v)=%v, wanted %v", req, resp, &tt.want)
		}

	}
}
