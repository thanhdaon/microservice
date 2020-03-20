package main

import (
	"fmt"
	"rabbitmq-tools/pkg/rabbit"
)

func main() {
	r := rabbit.SetupRabbit(
		"amqp://congtyio_email_crawler:FQ914bquqmkcW8N5aDhg6qzfIBDNLX8r@congty.io:5672/congtyio_email_crawler",
		[]string{},
	)

	consumers, err := r.FetchAllConsumer(
		"http://congty.io:15672/api/consumers",
		"congtyio_email_crawler",
		"FQ914bquqmkcW8N5aDhg6qzfIBDNLX8r",
	)

	if err != nil {
		fmt.Println(err)
	}

	for _, c := range consumers {
		fmt.Printf("%s-%s \n", c.Queue.Name, c.ConsumerTag)
	}
}
