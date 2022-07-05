package main

import (
	"context"
	"fmt"
	"grpc_study/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	address := "localhost:8080"
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	client := pb.NewGreetingServiceClient(conn)
	res, err := client.Hello(context.Background(), &pb.GreetingRequest{
		Name:   "Johnny",
		Age:    30,
		Gender: pb.GreetingRequest_MALE,
	})
	if err != nil {
		fmt.Println("server error:", err)
	}
	fmt.Println("server response:", res.Message)

	defer conn.Close()
}
