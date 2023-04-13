package interceptor

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func ClientUnaryLoggingInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	log.Println("=============================================")
	log.Println("Hello from client interceptor")
	log.Println("before invoker")
	log.Println("req = ", req)
	log.Println("----------")
	log.Println("invoker call:")
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Println("----------")
	log.Println("after invoker")
	log.Println("reply =", reply)
	log.Printf("invoked rpc method=%s; error=%v\n", method, err)
	grpclog.Infof("invoked rpc method=%s; error=%v\n", method, err)
	log.Println("=============================================")
	return err
}

func ClientStreamLoggingInterceptor(
	ctx context.Context,
	desc *grpc.StreamDesc,
	cc *grpc.ClientConn,
	method string,
	streamer grpc.Streamer,
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	log.Println("=============================================")
	log.Println("calling stream client interceptor for method", method)
	log.Println("=============================================")
	return streamer(ctx, desc, cc, method, opts...)
}
