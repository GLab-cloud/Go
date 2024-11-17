package main

import (
	pb "grpc/MicroServices/User"
	"log"
	"net"

	"context"

	grpc "google.golang.org/grpc"
)

type server struct {
	request []*pb.GetRequest
}
func (s *server) CreateUser(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{}, nil

}
func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterUserServer(srv, &server{})
	srv.Serve(lis)
}
