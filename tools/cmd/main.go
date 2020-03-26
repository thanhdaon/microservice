package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"tools/pkg/rabbit"
)

var (
	PROD_RABBIT    = "amqp://ltservices:jwtm9gby6Wn9VmvR@54.38.40.255:5672"
	STAGING_RABBIT = "amqp://admin:Mektoube2020@51.178.63.192:30002"
)

func main() {

	r := rabbit.SetupRabbit(STAGING_RABBIT, []string{})

	csvfile, err := os.Open("upload-photo.csv")
	failOnError(err, "can mot open file")
	reader := csv.NewReader(csvfile)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		failOnError(err, "okok")
		r.PublishToRabbit(record[0], "moderation-upload-photo")
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %+v", msg, err)
	}
}
