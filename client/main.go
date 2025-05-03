package main

import (
	"context"
	"fmt"
	pb "github.com/ankitgithubid/user-detail-grpc/github.com/ankitgithubid/grpc-server"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("Localhost:5051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	req := &pb.UserRequest{
		Id: "123",
	}

	resp, err := client.GetUser(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to get user: %v", err)
	}

	fmt.Println("UserResponse: ", resp)
}
