package main

import (
	"fmt"
	pb "grpc-adv/api/data"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, client := GetClient()
	defer conn.Close()

	fmt.Println("simple")
	err := GetUnary(client)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println()
	fmt.Println("server stream")
	err = GetServerStream(client)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println()
	fmt.Println("client stream:")
	err = GetClientStream(client)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println()
	fmt.Println("bidirectional stream:")
	err = GetBidirectional(client)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetClient() (*grpc.ClientConn, pb.DataClient) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(":8000", opts...)
	if err != nil {
		log.Fatalln("fail to dial:", err)
	}
	return conn, pb.NewDataClient(conn)
}
