package main

import (
	"context"
	"greet/greetpb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	failOnError(err, "Could not connect")
	defer connection.Close()

	client := greetpb.NewGreetServiceClient(connection)
	log.Printf("[INFO] Client created")

	// doUnary(client)
	// doServerStreaming(client)
	// doClientStreaming(client)
	doBidirectionalStreaming(client)

}

func doUnary(client greetpb.GreetServiceClient) {
	log.Println("[INFO] Starting to do a Unary RPC ...")
	req := &greetpb.GreetingRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Thanh",
			LastName:  "Dao Ngoc",
		},
	}
	res, err := client.Greet(context.Background(), req)
	failOnError(err, "err while calling Greet rpc")

	log.Printf("[INFO] Response from Greet: %v", res.Result)
}

func doServerStreaming(client greetpb.GreetServiceClient) {
	log.Println("[INFO] Start to do a Server Streaming RPC ...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Thanh",
			LastName:  "Dao Ngoc",
		},
	}
	stream, err := client.GreetManyTimes(context.Background(), req)
	failOnError(err, "Could not get stream")
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		failOnError(err, "error while reading stream")
		log.Printf("[INFO] Response from GreetmanyTimes: %v", msg.GetResult())
	}
}

func doClientStreaming(client greetpb.GreetServiceClient) {
	log.Println("Starting to do a Server Streaming RPC")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Thanh-1",
				LastName:  "Dao Ngoc",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Thanh-2",
				LastName:  "Dao Ngoc",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Thanh-3",
				LastName:  "Dao Ngoc",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Thanh-4",
				LastName:  "Dao Ngoc",
			},
		},
	}
	stream, err := client.LongGreet(context.Background())
	failOnError(err, "error while calling LongGreet")

	for _, req := range requests {
		log.Printf("Sending req: %v \n", req)
		stream.Send(req)
		time.Sleep(500 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	failOnError(err, "error while receiving response from LongGreet")
	log.Printf("LongGreet response: %v\n", res)
}

func doBidirectionalStreaming(client greetpb.GreetServiceClient) {
	log.Println("Starting to do a BiDi Streaming RPC ...")

	stream, err := client.GreetEveryone(context.Background())
	failOnError(err, "error while creating stream")

	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Thanh-1",
				LastName:  "Dao Ngoc",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Thanh-2",
				LastName:  "Dao Ngoc",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Thanh-3",
				LastName:  "Dao Ngoc",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Thanh-4",
				LastName:  "Dao Ngoc",
			},
		},
	}

	forever := make(chan bool)
	go func() {
		for _, req := range requests {
			log.Printf("Sending message: %s", req.GetGreeting().GetFirstName())
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			failOnError(err, "error while receiving")
			log.Printf("Received: %v", res.GetResult())
		}
		close(forever)

	}()
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("[ERROR] %s: %v", msg, err)
	}
}
