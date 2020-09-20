package interceptors

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func StreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	md, ok := metadata.FromIncomingContext(ss.Context())

	if !ok {
		return status.Errorf(codes.Internal, "Invalid Token or Internal error")
	}
	if _, ok := md["authorization"]; !ok {
		return status.Errorf(codes.Unauthenticated, "Token not provided using authorization metadata")
	}
	if t := md["authorization"][0]; t == "" {
		return status.Errorf(codes.Unauthenticated, "Token not provided using authorization metadata")
	}

	authorized, err := checkTokenInterceptor(md["authorization"][0])
	if err != nil {
		log.Println(err)
		return err
	}
	if !authorized {
		return status.Errorf(codes.Unauthenticated, "Invalid Token")
	}

	return handler(srv, ss)

}
