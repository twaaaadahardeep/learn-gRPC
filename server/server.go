package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/twaaaadahardeep/learn-gRPC/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8081, "")
)

type server struct {
	proto.UnimplementedMessageServer
}

func (s *server) GetMessage(ctx context.Context, in *proto.User) (*proto.UserMessage, error) {
	log.Printf("Received request from %v", in.GetName())

	return &proto.UserMessage{
		User:    in,
		Message: "Hello motherfucker!!!!",
	}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterMessageServer(s, &server{})

	log.Printf("Server listening at port: %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
