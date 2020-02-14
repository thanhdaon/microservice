package main

import (
	"context"
	"io"
	"log"
	"time"

	"calculator/pb"

	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	failOnError(err, "Could not connect to server")
	defer connection.Close()

	client := pb.NewCalculatorClient(connection)
	log.Println("[INFO] client created")
	// doUnary(client)
	// doServerStreaming(client)
	doClientStreaming(client)
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

func doClientStreaming(client pb.CalculatorClient) {
	log.Println("Starting to do a Client Streaming RPC")
	nums := []int32{1, 2, 3, 4, 5, 6}
	stream, err := client.ComputeAverage(context.Background())
	failOnError(err, "error while calling ComputeAverage")
	for _, num := range nums {
		log.Printf("Sending number : %d", num)
		stream.Send(&pb.ComputeAverageRequest{Num: num})
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	failOnError(err, "error while receiving response from ComputeAverage")
	log.Printf("ComputeAverage response: %v", res)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("[ERROR] %s: \n	%v", msg, err)
	}
}
