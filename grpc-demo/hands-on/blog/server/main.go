package main

import (
	"blog/pb"
	"context"
	"fmt"

	"log"
	"net"
	"os"
	"os/signal"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogResponse, error) {
	blog := req.GetBlog()
	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}

	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to OID"),
		)
	}

	return &pb.CreateBlogResponse{
		Blog: &pb.Blog{
			Id:       oid.Hex(),
			AuthorId: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		},
	}, nil
}

func (*server) ReadBlog(ctx context.Context, req *pb.ReadBlogRequest) (*pb.ReadBlogResponse, error) {
	log.Println("Read Blog Request")
	oid, err := primitive.ObjectIDFromHex(req.GetBlogId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}
	data := &blogItem{}
	filter := bson.D{{Key: "_id", Value: oid}}

	if err := collection.FindOne(context.Background(), filter).Decode(&data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("cannot find blog with ID: %v", err),
		)
	}
	return &pb.ReadBlogResponse{Blog: &pb.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorID,
		Content:  data.Content,
		Title:    data.Title,
	}}, nil
}

func main() {
	// if ea crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	connectDB()
	defer client.Disconnect(context.TODO())

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	failOnError(err, "Failed to listen")

	grpcServer := grpc.NewServer()
	pb.RegisterBlogServiceServer(grpcServer, &server{})

	go func() {
		log.Println("[INFO] Server listenning on port 50051")
		failOnError(grpcServer.Serve(listener), "Failed to server")
	}()

	// wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	grpcServer.Stop()
	listener.Close()
	log.Println("[INFO] The server stopped!")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("[ERROR] %s %v", msg, err)
	}
}
