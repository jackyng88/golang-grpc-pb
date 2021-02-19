package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"../greetpb"
)

func main() {

	fmt.Println("Hai, Client-desu.")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not establish connection: %v", err)
	}

	// Since we want the close() to come at the very end, the defer
	// allows us to have it happen at the end.
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created client: %f", c)
}
