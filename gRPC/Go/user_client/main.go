package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "rest-graphql-grpc/grpc/go/protos"
)

const (
	address     = "localhost:50051"
)


func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagerClient(conn)

	id := 1
	// if len(os.Args) > 1 {
	// 	id = os.Args[1]
	// }
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// TODO
	r, err := c.(ctx, &pb.HelloRequest{ID: id})
	if err != nil {
		log.Fatalf("could not : %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}