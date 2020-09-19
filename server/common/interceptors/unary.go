package interceptors

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func UnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Errorf(codes.Internal, "Invalid Token or Internal error")
	}
	if _, ok := md["authorization"]; !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Token not provided using authorization metadata")
	}
	if t := md["authorization"][0]; t == "" {
		return nil, status.Errorf(codes.Unauthenticated, "Token not provided using authorization metadata")
	}

	authorized, err := checkTokenInterceptor(md["authorization"][0])
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if !authorized {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid Token")
	}

	return handler(ctx, req)
}
