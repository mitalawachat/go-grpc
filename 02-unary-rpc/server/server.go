package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mitalawachat/go-grpc/02-unary-rpc/welcomepb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Welcome(ctx context.Context, request *welcomepb.WelcomeRequest) (*welcomepb.WelcomeResponse, error) {
	fmt.Printf("Welcome Request: %v \n", request)

	firstName := request.GetPerson().GetFirstName()
	lastName := request.GetPerson().GetLastName()

	result := "Hello, " + firstName + " " + lastName + "!"

	response := &welcomepb.WelcomeResponse{
		Result: result,
	}
	return response, nil
}

func main() {
	fmt.Println("Hello Gophers! I'm server!")

	listener, listenError := net.Listen("tcp", "0.0.0.0:50051")
	if listenError != nil {
		log.Fatalf("failed to listen: %v", listenError)
	}

	s := grpc.NewServer()
	welcomepb.RegisterWelcomeServiceServer(s, &server{})

	if serverError := s.Serve(listener); serverError != nil {
		log.Fatalf("failed to serve: %v", serverError)
	}
}
