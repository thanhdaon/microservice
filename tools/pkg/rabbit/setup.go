package rabbit

import (
	"github.com/streadway/amqp"
)

func SetupRabbit(uri string, queues []string) *Rabbit {
	var err error
	rabbitConn, err := amqp.Dial(uri)
	failOnError(err, "Failed to connect to RabbitMQ")

	channel, err := rabbitConn.Channel()
	failOnError(err, "Failed to open a channel")

	err = channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	for _, queue := range queues {
		_, err = channel.QueueDeclare(
			queue, // name
			true,  // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		failOnError(err, "Failed to declare a queue")
	}

	return &Rabbit{Connection: rabbitConn, Channel: channel, URI: uri}
}
