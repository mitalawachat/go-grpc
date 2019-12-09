package main

import (
	"fmt"
	"log"
	"net"

	"github.com/mitalawachat/go-grpc/01-code-generation/messagepb"
	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	fmt.Println("Hello Gophers! I'm server!")

	listener, listenError := net.Listen("tcp", "0.0.0.0:50051")
	if listenError != nil {
		log.Fatalf("failed to listen: %v", listenError)
	}

	s := grpc.NewServer()
	messagepb.RegisterMessageServiceServer(s, &server{})

	if serverError := s.Serve(listener); serverError != nil {
		log.Fatalf("failed to serve: %v", serverError)
	}
}
