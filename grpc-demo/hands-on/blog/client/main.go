package main

import (
	"blog/pb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	failOnError(err, "Could not connect")
	defer connection.Close()

	client := pb.NewBlogServiceClient(connection)
	log.Printf("[INFO] Client created")

	// create blog
	req := &pb.CreateBlogRequest{
		Blog: &pb.Blog{
			AuthorId: "ThanhDao",
			Title:    "HAHAHA",
			Content:  "Content of te firsr Blog",
		},
	}
	res, err1 := client.CreateBlog(context.Background(), req)
	failOnError(err1, "unexpected error")
	log.Printf("response : %v", res)

	// read blog
	_, err2 := client.ReadBlog(context.Background(), &pb.ReadBlogRequest{BlogId: "1111"})
	if err2 != nil {
		fmt.Printf("Error happened while reading: %v\n", err)
	}

	fmt.Println(res.GetBlog().GetId())
	readBlogReq := &pb.ReadBlogRequest{BlogId: res.GetBlog().GetId()}
	res2, err3 := client.ReadBlog(context.Background(), readBlogReq)
	if err3 != nil {
		fmt.Printf("Error happened while reading: %v", err)
	} else {
		fmt.Printf("ReadBlogResponse: %v", res2)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("[ERROR] %s: %v", msg, err)
	}
}
