package main

import (
	pb "grpc/MicroServices/User"
	"log"
	
	"context"

	grpc "google.golang.org/grpc"
)
func main() {
	conn, err := grpc.Dial("user:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)
    request := &pb.GetRequest{
		Keyword: "params request",
	}
	resp, err := client.CreateUser(context.Background(), request)
    // To do something with resp from instance server response
}
