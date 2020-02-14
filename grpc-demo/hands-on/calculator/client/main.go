package main

import (
	"context"
	"io"
	"log"

	"calculator/pb"

	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	failOnError(err, "Could not connect to server")
	defer connection.Close()

	client := pb.NewCalculatorClient(connection)
	log.Println("[INFO] client created")
	doUnary(client)
	doServerStreaming(client)
}

func doUnary(client pb.CalculatorClient) {
	log.Println("[INFO] Starting to do a Unary RPC ...")
	req := &pb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 20,
	}
	res, err := client.Sum(context.Background(), req)
	failOnError(err, "err while call Sum RPC")

	log.Printf("[INFO] Response from Sum: %v", res.Result)
}

func doServerStreaming(client pb.CalculatorClient) {
	log.Println("[INFO] Starting to do a Server Streaming RPC (Number=120)")
	req := &pb.PrimeDecomposeRequest{Number: 120}
	stream, err := client.PrimeDecompose(context.Background(), req)
	failOnError(err, "Could not get stream")
	for {
		decomposedNumber, err := stream.Recv()
		if err == io.EOF {
			break
		}
		failOnError(err, "error while reading stream")
		log.Printf("[INFO] Response from PrimeDecompose: %v", decomposedNumber.GetResult())
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("[ERROR] %s: \n	%v", msg, err)
	}
}
