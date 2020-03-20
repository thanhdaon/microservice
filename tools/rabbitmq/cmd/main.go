package main

import (
	"fmt"
	"log"
	"rabbitmq-tools/pkg"
)

func main() {
	consumers, err := pkg.FetchAllConsumer(
		"http://51.178.63.192:30001/api/consumers",
		"admin",
		"Mektoube2020",
	)

	failOnError(err, "main")
	for _, consumer := range consumers {
		if consumer.Queue.Name == "moderation-image-queue" {
			fmt.Printf("%s - %s\n", consumer.Queue.Name, consumer.ConsumerTag)
		}
	}

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}
