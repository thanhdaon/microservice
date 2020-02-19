package crawler

import (
	"log"

	"github.com/streadway/amqp"
)

var rabbitConn *amqp.Connection
var channel *amqp.Channel

func SetupRabbit() {
	var err error

	rabbitConn, err = amqp.Dial("amqp://guest:guest@localhost:5672")
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
		"emailsvc-domain-waiting-to-crawl", // name
		true,                               // durable
		false,                              // delete when unused
		false,                              // exclusive
		false,                              // no-wait
		nil,                                // arguments
	)
	failOnError(err, "Failed to declare a queue")
}

func CleanupRabbit() {
	rabbitConn.Close()
	channel.Close()
}

func Start() {
	msgs, err := channel.Consume(
		"emailsvc-domain-waiting-to-crawl", // queue
		"",                                 // consumer
		false,                              // auto-ack
		false,                              // exclusive
		false,                              // no-local
		false,                              // no-wait
		nil,                                // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for msg := range msgs {
			crawl(string(msg.Body))
			msg.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for queue: emailsvc-domain-waiting-to-crawl")
}

func publishToRabbit(body string) {
	err := channel.Publish(
		"",                                 // exchange
		"emailsvc-domain-waiting-to-crawl", // routing key
		false,                              // mandatory
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
