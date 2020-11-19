package common_testing

import (
	"context"

	evgrpc "github.com/ktr0731/evans/grpc"
)

// ContextGenerator aims to create a testing context (ONLY FOR DEV AND TEST MODE)
// ctx is the context, it can contain attached metadatas
// service is the name of the service, but the value has often the format "package.Service" => auth.AuthService for example.
// method is the name of the method.
// port is the port at which the services is exposed.
// req is the request, but declared like req := &authpb.CreateNewUserRequest{Username:..., Email: ...} the & is really important
// res is the response, must be like res := &authpb.CreateNewUserResponse{}, the & is really important.
// The res variable will be modified by reference !
func ContextGenerator(ctx context.Context, service, method, port string, req interface{}, res interface{}) (interface{}, error) {

	cli, err := evgrpc.NewClient("server:"+port, "server", true, false, "", "", "")
	if err != nil {
		return nil, err
	}

	// req := &authpb.CreateNewUserRequest{
	// 	Email:     "toto",
	// 	FirstName: "toto",
	// 	LastName:  "toto",
	// 	Username:  "toto",
	// 	Password:  "toto",
	// }

	// res := &authpb.CreateNewUserResponse{}

	_, _, err = cli.Invoke(ctx, service+"."+method, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
