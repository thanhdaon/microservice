package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id`
	Content  string             `bson:"content`
	Title    string             `bson:"title"`
}

func connectDB() {
	var err error
	ctx := context.TODO()
	opts := options.Client().ApplyURI("mongodb://root:root_password@localhost:27017")
	client, err = mongo.Connect(ctx, opts)
	failOnError(err, "could not connect to mongo")
	collection = client.Database("demo").Collection("blog")
	log.Println("[INFO] mongo connected")
}
