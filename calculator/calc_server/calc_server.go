package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/igorvirgilio/grpc-go-course/calculator/calcpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calculator(ctx context.Context, req *calcpb.CalculatorRequest) (*calcpb.CalculatorReponse, error) {
	fmt.Printf("Calculation function was invoked with %v\n", req)
	number1 := req.GetCalculating().GetNumber1()
	number2 := req.GetCalculating().GetNumber2()
	sum := number1 + number2
	res := &calcpb.CalculatorReponse{
		Result: sum,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
