package common_testing

import (
	"context"
	common_kvdb "primitivofr/owly/server/common/kvdb"

	"google.golang.org/grpc/metadata"
)

// PrepareCtxWithToken returns a ctx with an embedded authorization token.
// First one is Incoming, second is OutGoing.
func PrepareCtxWithToken(username string) (context.Context, context.Context) {

	db, err := common_kvdb.InitDB("testingdb", "tokens")
	CheckErr(err, "")

	accessToken, readallErr := common_kvdb.ReadData(db, "tokens", username)
	CheckErr(readallErr, "Error while reading token for "+username)

	db.Close()

	md := metadata.Pairs("authorization", string(accessToken))
	return metadata.NewIncomingContext(context.Background(), md), metadata.NewOutgoingContext(context.Background(), md)

}
