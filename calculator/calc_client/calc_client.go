package main

import (
	"context"
	"fmt"
	"log"

	"github.com/igorvirgilio/grpc-go-course/calculator/calcpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I want to calulate")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := calcpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(c calcpb.CalculatorServiceClient) {
	req := &calcpb.CalculatorRequest{
		Calculating: &calcpb.Calculating{
			Number1: 3,
			Number2: 10,
		},
	}
	res, err := c.Calculator(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calculator RPC: %v", err)
	}
	log.Printf("Response from Calculator is: %v", res.Result)
}
