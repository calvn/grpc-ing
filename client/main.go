package main

import (
	"log"
	"os"

	pb "github.com/cleung2010/grpc-ing/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address        = "localhost:50051"
	defaultMessage = "Calvin"
)

func main() {
	// Connecting to the service
	//
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Calling a RPC
	message := defaultMessage
	if len(os.Args) > 1 {
		message = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Message: message})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
