package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mitalawachat/go-grpc/02-unary-rpc/welcomepb"
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

	welcomeRequest := &welcomepb.WelcomeRequest{
		Person: &welcomepb.Person{
			FirstName: "John",
			LastName:  "Wick",
		},
	}

	welcomeResponse, welcomeError := client.Welcome(context.Background(), welcomeRequest)
	if welcomeError != nil {
		log.Fatalf("failed to call welcome service: %v", welcomeError)
	}

	fmt.Printf("Welcome Response: %v \n", welcomeResponse)
}
