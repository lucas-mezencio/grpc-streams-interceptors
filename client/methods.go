package main

import (
	"context"
	"fmt"
	pb "grpc-adv/api/data"
	"io"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetUnary(client pb.DataClient) error {
	res, err := client.GetByID(context.Background(), &pb.Request{
		RequestAt: timestamppb.Now(),
		ID:        "X1",
	})
	if err != nil {
		return fmt.Errorf("error single: %w", err)
	}
	fmt.Println(res.ResponseAt.AsTime().Format(time.RFC3339), ":", res.Data)
	return nil
}

func GetServerStream(client pb.DataClient) error {
	stream, err := client.GetAll(context.Background(), &pb.Request{
		RequestAt: timestamppb.Now(),
		ID:        "SrvStream",
	})
	if err != nil {
		return fmt.Errorf("error server stream: %w", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetAll err: %v", client, err)
		}
		fmt.Println(res.ResponseAt.AsTime().Format(time.RFC3339), ":", res.Data)
	}
	return nil
}

func GetClientStream(client pb.DataClient) error {
	stream, err := client.SendAll(context.Background())
	if err != nil {
		return fmt.Errorf("error client stream: %w", err)
	}
	for i := 0; i < 10; i++ {
		if err := stream.Send(&pb.Request{
			RequestAt: timestamppb.Now(),
			ID:        fmt.Sprintf("clientstream#%d", i),
		}); err != nil {
			return fmt.Errorf("client stream err: %w", err)
		}
		time.Sleep(time.Second)
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("%v.CloseAndRecv() got error %w", stream, err)
	}
	log.Println(reply)
	return nil
}
