package main

import (
	"fmt"
	"grpc_study/pb"
	"grpc_study/service"

	"net"

	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	pb.RegisterGreetingServiceServer(server, service.GreetingService{})
	port := ":8080"
	listener, _ := net.Listen("tcp", port)
	fmt.Println("grpc server running at", port)
	server.Serve(listener)
}
