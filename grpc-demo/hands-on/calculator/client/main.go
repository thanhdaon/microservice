package main

import (
	"context"
	"io"
	"log"
	"time"

	"calculator/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	failOnError(err, "Could not connect to server")
	defer connection.Close()

	client := pb.NewCalculatorClient(connection)
	log.Println("[INFO] client created")
	// doUnary(client)
	// doServerStreaming(client)
	// doClientStreaming(client)
	// doBiDiStreaming(client)
	doErrorUnary(client)
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

func doBiDiStreaming(client pb.CalculatorClient) {
	log.Println("Starting to do a BiDi client Streaming RPC")
	nums := []int32{1, 5, 3, 6, 2, 20}
	stream, err := client.FindMax(context.Background())
	failOnError(err, "error while calling FindMax")

	waitC := make(chan bool)
	go func() {
		for _, num := range nums {
			log.Printf("send num = %d\n", num)
			stream.Send(&pb.FindMaxRequest{Num: num})
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
			failOnError(err, "error while receiving response")
			log.Printf("Max = %d\n", res.GetMax())
		}
		close(waitC)
	}()
	<-waitC
}

func doErrorUnary(client pb.CalculatorClient) {
	log.Printf("Starting to do a SquareRoot RPC...")
	req := &pb.SquareRootRequest{Number: -10}
	res, err := client.SquareRoot(context.Background(), req)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error from GRPC (user error)
			log.Println(respErr.Message())
			return
		} else {
			// framework error
			failOnError(err, "Big error calling SquareRoot")
		}
	}
	log.Printf("Result of square root of %v: %v\n", 10, res.GetNumberRoot())
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("[ERROR] %s: \n	%v", msg, err)
	}
}
