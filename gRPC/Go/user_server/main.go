package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "rest-graphql-grpc/grpc/go/protos"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedUserManagerServer
}

func (s *server) getUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("Received")
	user := struct {
		id   int
		name string
	}{
		1,
		"foo",
	}
	// TODO
	return &pb.UserResponse{user}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
