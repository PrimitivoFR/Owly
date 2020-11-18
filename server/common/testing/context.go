package common_testing

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// PrepareCtxWithToken returns a ctx with an embedded authorization token.
// First one is Incoming, second is OutGoing.
func PrepareCtxWithToken(username string) (context.Context, context.Context) {

	accessToken := LoadFromKvdb("tokens", username)

	md := metadata.Pairs("authorization", string(accessToken))
	return metadata.NewIncomingContext(context.Background(), md), metadata.NewOutgoingContext(context.Background(), md)

}
