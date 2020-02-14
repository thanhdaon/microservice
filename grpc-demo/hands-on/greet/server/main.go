package main

import (
	"context"
	"log"
	"net"

	"greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetingRequest) (*greetpb.GreetingResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	result := "hello " + firstName
	return &greetpb.GreetingResponse{Result: result}, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	failOnError(err, "Failed to listen")

	grpcServer := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(grpcServer, &server{})

	log.Println("Server listenning on port 50051")
	failOnError(grpcServer.Serve(listener), "Failed to server")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("[ERROR] %s %v", msg, err)
	}
}
