package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/mitalawachat/go-grpc/05-bi-directional-streaming-rpc/welcomepb"
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

	stream, streamError := client.WelcomeEveryone(context.Background())
	if streamError != nil {
		log.Fatalf("failed to call service: %v", streamError)
	}

	waitChannel := make(chan struct{})

	go sendMessages(stream)
	go receiveMessages(stream, waitChannel)

	<-waitChannel
}

func sendMessages(stream welcomepb.WelcomeService_WelcomeEveryoneClient) {
	requests := []*welcomepb.WelcomeEveryoneRequest{
		{
			Person: &welcomepb.Person{
				FirstName: "Steve",
				LastName:  "Rogers",
			},
		},
		{
			Person: &welcomepb.Person{
				FirstName: "Tony",
				LastName:  "Stark",
			},
		},
		{
			Person: &welcomepb.Person{
				FirstName: "Natalia",
				LastName:  "Romanova",
			},
		},
	}
	for i, request := range requests {
		sendError := stream.Send(request)
		if sendError != nil {
			log.Fatalf("failed to send data: %v", sendError)
		}

		fmt.Printf("Sent %d message \n", i+1)
		time.Sleep(1000 * time.Millisecond)
	}
	stream.CloseSend()
}

func receiveMessages(stream welcomepb.WelcomeService_WelcomeEveryoneClient, waitChannel chan struct{}) {
	for {
		message, messageError := stream.Recv()
		if messageError == io.EOF {
			close(waitChannel)
			break
		}
		if messageError != nil {
			log.Fatalf("error while reading stream: %v", messageError)
		}
		fmt.Printf("Received Response: %v \n", message.GetResult())
	}
}
