package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/mitalawachat/go-grpc/03-server-streaming-rpc/welcomepb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) WelcomeManyTimes(request *welcomepb.WelcomeManyTimesRequest, stream welcomepb.WelcomeService_WelcomeManyTimesServer) error {
	fmt.Printf("Received Request: %v \n", request)

	firstName := request.GetPerson().GetFirstName()
	lastName := request.GetPerson().GetLastName()

	for i := 0; i <= 10; i++ {
		result := fmt.Sprintf("[%d] Hello, %s %s!", i, firstName, lastName)
		response := &welcomepb.WelcomeManyTimesResponse{
			Result: result,
		}
		stream.Send(response)
		time.Sleep(1000 * time.Millisecond)
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
