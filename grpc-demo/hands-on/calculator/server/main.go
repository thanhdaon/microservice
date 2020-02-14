package main

import (
	"calculator/pb"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type service struct{}

func (*service) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()

	return &pb.SumResponse{Result: firstNumber + secondNumber}, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	failOnError(err, "Failed to listen")

	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServer(grpcServer, &service{})

	log.Println("Server listenning on port 50051")
	failOnError(grpcServer.Serve(listener), "Failed to server")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("[ERROR] %s %v", msg, err)
	}
}
