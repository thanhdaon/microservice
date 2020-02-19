import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	connectDB()
	update()
	defer client.Disconnect(context.TODO())
}

func connectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://root:root_password@localhost:27017")

	var err error

	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func insert() {
	collection := client.Database("demo").Collection("person")

	ruan := Person{"Ruan", 34, "Cape Town"}
	james := Person{"James", 32, "Nairobi"}
	frankie := Person{"Frankie", 31, "Nairobi"}

	trainers := []interface{}{ruan, james, frankie}
	insertResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a SINGLE document: ", insertResult.InsertedIDs)
}

func update() {
	collection := client.Database("demo").Collection("person")
	filter := bson.D{}
	update := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "age", Value: 1},
		}},
	}

	updateResult, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("update a SINGLE document", updateResult.MatchedCount, updateResult.ModifiedCount)
}
