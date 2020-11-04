package commontesting

import (
	"context"

	//evgrpc "github.com/ktr0731/evans/grpc"
	"google.golang.org/grpc"
	//rpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

type server struct{}

type ServiceClientFunc func(grpc.ClientConnInterface) interface{}

func ContextGenerator(sClient ServiceClientFunc, methodName string, in *struct{}, ctx context.Context) {

	// conn, _ := grpc.Dial("server:50054", grpc.WithInsecure())

	// cli := grpcreflect.NewClient(context.Background(), rpb.NewServerReflectionClient(conn))

	// b, _ := cli.ResolveService("AuthService")
	// evgrpc.NewClient()

	// c := b.GetMethods()
	// for _, kk := range c {
	// 	kk.AsProto
	// }
	// c := sClient(conn)
	// ret := reflect.ValueOf(c).MethodByName(methodName).Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(in)})

	// fmt.Println(ret)
}
