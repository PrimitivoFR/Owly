package common_jwt

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func ReadUUIDFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("Error while reading metadatas.")
		return "", status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while reading metadatas."),
		)
	}
	user_id, err := ExtractUUIDfromJWT(md["authorization"][0])
	if err != nil {
		log.Printf("Error while decoding jwt %v", err)
		return "", status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while decoding jwt %v", err),
		)
	}

	return user_id, nil
}
