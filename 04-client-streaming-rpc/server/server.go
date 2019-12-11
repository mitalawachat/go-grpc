package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/mitalawachat/go-grpc/04-client-streaming-rpc/welcomepb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) LongWelcome(stream welcomepb.WelcomeService_LongWelcomeServer) error {

	for {
		message, messageError := stream.Recv()

		if messageError == io.EOF {
			result := "Hello Gopher! We've received all your messages!"
			response := &welcomepb.LongWelcomeResponse{
				Result: result,
			}
			sendError := stream.SendAndClose(response)
			if sendError != nil {
				log.Fatalf("failed to send response: %v", sendError)
			}
			break
		}

		if messageError != nil {
			log.Fatalf("error while reading stream: %v", messageError)
		}

		fmt.Printf("Received Message: %v \n", message)
	}

	return nil
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
