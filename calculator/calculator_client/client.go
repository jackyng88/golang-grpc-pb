package main

import (
	"context"
	"fmt"
	"log"

	"../calculatorpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hai, Calculator Client-desu.")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not establish connection: %v", err)
	}

	// Since we want the close() to come at the very end, the defer
	// allows us to have it happen at the end.
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	//fmt.Printf("Created client: %f", c)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Sum Unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Calculator: %v", res.SumResult)
}
