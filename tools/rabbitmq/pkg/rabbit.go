package pkg

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

var (
	rabbitConn *amqp.Connection
	channel    *amqp.Channel
)

const (
	BING_SEARCH_RESULT_QUEUE = "emailsvc-bing-search-result"
	JS_BASED_WEBSITE_QUEUE   = "emailsvc-js-based-website"
)

func SetupRabbit(rabbitURI string) {
	var err error
	rabbitConn, err = amqp.Dial(rabbitURI)
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
		BING_SEARCH_RESULT_QUEUE, // name
		true,                     // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		nil,                      // arguments
	)
	failOnError(err, "Failed to declare a queue")
}

func CleanupRabbit() {
	rabbitConn.Close()
	channel.Close()
}

func publishToRabbit(body, queue string) {
	err := channel.Publish(
		"",    // exchange
		queue, // routing key
		false, // mandatory
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

type ConsumerDetail struct {
	ConsumerTag string `json:"consumer_tag"`
	Queue       struct {
		Name  string
		Vhost string
	}
}

func FetchAllConsumer(uri, username, password string) ([]ConsumerDetail, error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, "fetchAllConsumer")
	}
	req.Header.Add("Authorization", "Basic "+basicAuth(username, password))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "fetchAllConsumer")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "fetchAllConsumer")
	}

	var consumers []ConsumerDetail
	if err = json.Unmarshal(body, &consumers); err != nil {
		return nil, errors.Wrap(err, "fetchAllConsumer")
	}

	return consumers, nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}
