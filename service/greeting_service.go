package service

import (
	"context"
	"fmt"
	"grpc_study/pb"
)

type GreetingService struct {
	pb.UnimplementedGreetingServiceServer
}

func (g GreetingService) Hello(ctx context.Context, request *pb.GreetingRequest) (*pb.Greeting, error) {
	return &pb.Greeting{
		Message: fmt.Sprintf("Hello %s, you are %d years old and your gender is: %s", request.Name, request.Age, request.Gender.String()),
	}, nil
}
