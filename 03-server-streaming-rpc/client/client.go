package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/mitalawachat/go-grpc/03-server-streaming-rpc/welcomepb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello Gophers! I'm client!")

	connection, connectionError := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if connectionError != nil {
		log.Fatalf("could not connnect: %v", connectionError)
	}
	defer connection.Close()

	client := welcomepb.NewWelcomeServiceClient(connection)

	welcomeRequest := &welcomepb.WelcomeManyTimesRequest{
		Person: &welcomepb.Person{
			FirstName: "John",
			LastName:  "Wick",
		},
	}

	stream, welcomeError := client.WelcomeManyTimes(context.Background(), welcomeRequest)
	if welcomeError != nil {
		log.Fatalf("failed to call welcome service: %v", welcomeError)
	}

	for {
		message, messageError := stream.Recv()
		if messageError == io.EOF {
			break
		}
		if messageError != nil {
			log.Fatalf("error while reading welcome stream: %v", messageError)
		}
		fmt.Printf("Welcome Response: %v \n", message.GetResult())
	}
}
