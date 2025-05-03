package main

import (
	"context"
	"fmt"
	pb "github.com/ankitgithubid/user-detail-grpc/github.com/ankitgithubid/grpc-server"
	"google.golang.org/grpc"
	"log"
	"net"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *userServiceServer) GetUser(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	testUser := map[string]*pb.UserResponse{
		"123": {
			Id:    "123",
			Name:  "John",
			Email: "john@test.com",
		},
		"2": {
			Id:    "2",
			Name:  "Alice",
			Email: "alice@test.com",
		},
	}

	if user, ok := testUser[request.Id]; ok {
		return user, nil
	}

	return nil, fmt.Errorf("User ID :%s not found", request.GetId())
}

func main() {
	lis, err := net.Listen("tcp", ":5051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &userServiceServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
