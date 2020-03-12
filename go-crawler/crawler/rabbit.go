package crawler

import (
	"log"

	"github.com/streadway/amqp"
)

var rabbitConn *amqp.Connection
var channel *amqp.Channel

func SetupRabbit() {
	var err error
	rabbitConn, err = amqp.Dial("amqp://congtyio_email_crawler:FQ914bquqmkcW8N5aDhg6qzfIBDNLX8r@congty.io:5672/congtyio_email_crawler")
	failOnError(err, "Failed to connect to RabbitMQ")

	channel, err = rabbitConn.Channel()
	failOnError(err, "Failed to open a channel")

	err = channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	_, err = channel.QueueDeclare(
		"emailsvc-bing-search-result", // name
		true,                            // durable
		false,                           // delete when unused
		false,                           // exclusive
		false,                           // no-wait
		nil,                             // arguments
	)
	failOnError(err, "Failed to declare a queue")
}

func CleanupRabbit() {
	rabbitConn.Close()
	channel.Close()
}

func publishToRabbit(body string) {
	err := channel.Publish(
		"",                              // exchange
		"emailsvc-bing-search-result", // routing key
		false,                           // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		},
	)
	if err != nil {
		log.Printf("Failed to publish! \n%v\n", err)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
