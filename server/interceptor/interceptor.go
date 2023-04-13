package interceptor

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		log.Println("=============================================")
		log.Println("called unary server interceptor")
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, fmt.Errorf("couldn't parse incoming context metadata")
		}
		fmt.Println(md)

		r, err := handler(ctx, req)
		fmt.Printf("req = %v ;; res = %v\n", req, r)
		log.Println("=============================================")
		return r, err
	}
}

func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		fmt.Println("=============================================")
		log.Println("call server stream interceptor")
		log.Println("method = ", info.FullMethod)
		log.Println("is server stream = ", info.IsServerStream)
		log.Println("is client stream = ", info.IsClientStream)
		fmt.Println("=============================================")
		return handler(srv, ss)
	}
}
