package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/thyagofr/grpc-go-unary/calcproto"
)

type server struct{}

func (s *server) Calculate(ctx context.Context, req *calcproto.CalcRequest) (*calcproto.CalcResponse, error) {
	number1 := req.GetCalc().GetNumber1()
	number2 := req.GetCalc().GetNumber2()
	res := &calcproto.CalcResponse{
		Result: (number1 + number2),
	}
	return res, nil
}

func main() {
	log.Println("Running server...")
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}
	serv := grpc.NewServer()
	calcproto.RegisterCalcServiceServer(serv, &server{})
	if err := serv.Serve(listener); err != nil {
		log.Fatalf("Failer to start grpc server : %v", err)
	}
	log.Println("Running server...")
}
