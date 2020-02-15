package main

import (
	"calculator/pb"
	"context"
	"io"
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

func (*service) PrimeDecompose(req *pb.PrimeDecomposeRequest, stream pb.Calculator_PrimeDecomposeServer) error {
	var k int32 = 2
	num := req.GetNumber()
	for num > 1 {
		if num%k == 0 {
			stream.Send(&pb.PrimeDecomposeResponse{Result: k})
			num = num / k
			continue
		}
		k = k + 1
	}
	return nil
}

func (*service) ComputeAverage(stream pb.Calculator_ComputeAverageServer) error {
	log.Println("[INFO] ComputeAverate funtion was involed")
	var nums []int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.ComputeAverageResponse{
				Result: getAvg(nums),
			})
		}
		failOnError(err, "err while reading client stream")
		nums = append(nums, req.GetNum())
	}
}

func (*service) FindMax(stream pb.Calculator_FindMaxServer) error {
	log.Println("FindMax function was invoked")
	var nums []int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		failOnError(err, "error while reading client streaming")
		nums = append(nums, req.GetNum())
		sendErr := stream.Send(&pb.FindMaxResponse{Max: getMax(nums)})
		failOnError(sendErr, "error when sending data to client")
	}
}

func getMax(nums []int32) int32 {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func getAvg(nums []int32) float32 {
	var sum int32 = 0
	var numCount int = 0
	for _, num := range nums {
		sum = sum + num
		numCount++
	}
	return float32(float32(sum) / float32(numCount))
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
