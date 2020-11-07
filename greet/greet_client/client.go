package main

import (
	"context"
	"fmt"
	"log"

	"github.com/igorvirgilio/grpc-go-course/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // use without SSL
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close() // defer will execute in the very end

	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Created Client: %f", c)

	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Igor",
			LastName:  "Araujo",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)

}
