package main

import (
	"context"
	"flag"
	pb "github.com/brunollopes/gRPC/helloworld"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	defaultName = "Hello, I am sending this message via gRPC, hope you like it."
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Message: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Got: %s", r.GetAnswer())
}
