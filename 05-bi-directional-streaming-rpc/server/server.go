package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/mitalawachat/go-grpc/05-bi-directional-streaming-rpc/welcomepb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) WelcomeEveryone(stream welcomepb.WelcomeService_WelcomeEveryoneServer) error {

	for {
		message, messageError := stream.Recv()

		if messageError == io.EOF {
			return nil
		}
		if messageError != nil {
			log.Fatalf("error while reading stream: %v", messageError)
			return messageError
		}

		fmt.Printf("Received Message: %v \n", message)

		result := "Hello " + message.GetPerson().GetFirstName() + "! We've received your message!"
		response := &welcomepb.WelcomeEveryoneResponse{
			Result: result,
		}
		sendError := stream.Send(response)
		if sendError != nil {
			log.Fatalf("failed to send response to client: %v", sendError)
		}
	}

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
