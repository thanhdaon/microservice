package main

import (
	"context"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetingRequest) (*greetpb.GreetingResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	result := "hello " + firstName
	return &greetpb.GreetingResponse{Result: result}, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		res := &greetpb.GreetManyTimesResponse{
			Result: "Hello " + firstName + strconv.Itoa(i),
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked with")
	result := "Hello "
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		failOnError(err, "error while reading client stream")
		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
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
