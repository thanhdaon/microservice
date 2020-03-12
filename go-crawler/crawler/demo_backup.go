package crawler

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func DemoPackup() {
	SetupRabbit()
	defer CleanupRabbit()

	file, err := os.OpenFile("backup.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	failOnError(err, "cannot create csv file")
	defer file.Close()

	writer := csv.NewWriter(file)

	msgs, err := channel.Consume(
		"demo-backup", // queue
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			writer.Write([]string{string(msg.Body)})
			writer.Flush()
			fmt.Println(string(msg.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
