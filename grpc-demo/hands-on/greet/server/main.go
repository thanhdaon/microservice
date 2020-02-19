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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
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

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone function involed with a streaming request")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		failOnError(err, "error while reading client streaming")
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! "
		sendErr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		failOnError(sendErr, "error while sending data tp client")
	}
}

func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
	log.Printf("GreetWithDeadline function was invoked!")
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			log.Println("the client canceled the request!")
			return nil, status.Error(
				codes.DeadlineExceeded,
				"The client canceled the request",
			)
		}
		time.Sleep(1 * time.Second)
	}
	firstName := req.GetGreeting().GetFirstName()
	result := "hello " + firstName
	return &greetpb.GreetWithDeadlineResponse{Result: result}, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	failOnError(err, "Failed to listen")

	opts := []grpc.ServerOption{}
	tls := false
	if tls {
		certFile := "../ssl/server.crt"
		keyFile := "../ssl/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
		failOnError(sslErr, "fail to loading certificates!")
		opts = append(opts, grpc.Creds(creds))
	}
	grpcServer := grpc.NewServer(opts...)
	greetpb.RegisterGreetServiceServer(grpcServer, &server{})

	log.Println("Server listenning on port 50051")
	failOnError(grpcServer.Serve(listener), "Failed to server")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("[ERROR] %s %v", msg, err)
		return
	}
}
