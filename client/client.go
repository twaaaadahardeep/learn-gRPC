package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/twaaaadahardeep/learn-gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:8081", "")
)

func main() {

	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewMessageClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	m, err := c.GetMessage(ctx, &proto.User{UserId: 1, Name: "Paaji"})
	if err != nil {
		log.Fatalf("Could not send Message: %v", err)
	}

	log.Printf("Message Received from %v : %v", m.GetUser(), m.GetMessage())
}
