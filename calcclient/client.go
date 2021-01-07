package main

import (
	"context"
	"fmt"
	"log"

	"github.com/thyagofr/grpc-go-unary/calcproto"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Running grpc client...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error creating grpc client : %v", err)
	}
	defer cc.Close()
	client := calcproto.NewCalcServiceClient(cc)
	response, err := client.Calculate(context.Background(), &calcproto.CalcRequest{
		Calc: &calcproto.Calc{
			Number1: 10,
			Number2: 3,
		},
	})
	if err != nil {
		log.Fatalf("An error occurred : %v", err)
	}
	fmt.Println(response.Result)
}
