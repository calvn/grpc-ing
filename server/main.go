package main

import (
	"log"
	"net"

	pb "github.com/cleung2010/grpc-ing/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Generic struct to implement helloworld.GreeterServer
type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: in.Message}, nil
}

const (
	port = ":50051"
)

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(listener)
	log.Println("Starting gRPC server...")
}
