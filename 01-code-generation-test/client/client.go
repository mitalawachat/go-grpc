package main

import (
	"fmt"
	"log"

	"github.com/mitalawachat/go-grpc/01-code-generation-test/messagepb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello Gophers! I'm client!")

	connection, connectionError := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if connectionError != nil {
		log.Fatalf("could not connnect: %v", connectionError)
	}
	defer connection.Close()

	client := messagepb.NewMessageServiceClient(connection)
	fmt.Printf("created client: %f", client)
}
