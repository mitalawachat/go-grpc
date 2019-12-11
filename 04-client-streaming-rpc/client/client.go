package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mitalawachat/go-grpc/04-client-streaming-rpc/welcomepb"
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

	stream, streamError := client.LongWelcome(context.Background())
	if streamError != nil {
		log.Fatalf("failed to call service: %v", streamError)
	}

	for i := 1; i <= 10; i++ {
		welcomeRequest := &welcomepb.LongWelcomeRequest{
			Person: &welcomepb.Person{
				FirstName: "John",
				LastName:  "Wick",
			},
		}

		sendError := stream.Send(welcomeRequest)
		if sendError != nil {
			log.Fatalf("failed to send data: %v", sendError)
		}

		fmt.Printf("Sent %d message \n", i)
		time.Sleep(1000 * time.Millisecond)
	}

	message, messageError := stream.CloseAndRecv()
	if messageError != nil {
		log.Fatalf("failed to receive data: %v", messageError)
	}

	fmt.Printf("Received Response: %v \n", message.GetResult())
}
