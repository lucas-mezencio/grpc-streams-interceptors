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

func (s *DataServer) GetByID(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	res := &pb.Response{
		ResponseAt: timestamppb.Now(),
		Data:       fmt.Sprintf("hello from response req(%v) at %v!", req.ID, req.RequestAt.AsTime().Format(time.RFC3339)),
	}
	return res, nil
}
func (s *DataServer) GetAll(req *pb.Request, stream pb.Data_GetAllServer) error {
	for i := 0; i < 10; i++ {
		if err := stream.Send(&pb.Response{
			ResponseAt: timestamppb.Now(),
			Data:       fmt.Sprintf("#%d hello from response req(%v) at %v!", i, req.ID, req.RequestAt.AsTime().Format(time.RFC3339)),
		}); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (s *DataServer) SendAll(stream pb.Data_SendAllServer) error {
	start := time.Now()
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Response{
				ResponseAt: timestamppb.Now(),
				Data:       fmt.Sprintf("finish reading data in %v", time.Since(start)),
			})
		}
		if err != nil {
			log.Println("error:", err)
			return err
		}
		log.Printf("hello from reqeust (%v) at %v!", req.ID, req.RequestAt.AsTime().Format(time.RFC3339))
	}
}

func (s *DataServer) SandAndGetAll(stream pb.Data_SandAndGetAllServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Println("error:", err)
			return err
		}
		time.Sleep(500 * time.Millisecond)
		log.Printf("hello from response (%v) at %v!", req.ID, req.RequestAt.AsTime().Format(time.RFC3339))
	}
}
