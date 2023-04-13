package main

import (
	pb "grpc-adv/api/data"
	"grpc-adv/server/interceptor"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var _ pb.DataServer = (*DataServer)(nil)

type DataServer struct {
	pb.UnimplementedDataServer
}

func (s *DataServer) Run() error {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println("error running server")
		return err
	}
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.UnaryServerInterceptor()),
		grpc.StreamInterceptor(interceptor.StreamServerInterceptor()),
	)
	pb.RegisterDataServer(srv, s)
	log.Println("server listening at", lis.Addr())

	reflection.Register(srv)
	return srv.Serve(lis)
}
