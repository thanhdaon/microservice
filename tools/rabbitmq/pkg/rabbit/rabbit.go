package rabbit

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

type ConsumerDetail struct {
	ConsumerTag string `json:"consumer_tag"`
	Queue       struct {
		Name  string
		Vhost string
	}
}

type Rabbit struct {
	URI        string
	Email      string
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func (r *Rabbit) PublishToRabbit(body, queue string) {
	err := r.Channel.Publish(
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

func (r *Rabbit) Consume(queue string) <-chan amqp.Delivery {
	msgChannel, err := r.Channel.Consume(
		queue, // queue
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	failOnError(err, "Failed to register a consumer")
	return msgChannel
}

func (r *Rabbit) FetchAllConsumer(url, username, password string) ([]ConsumerDetail, error) {
	req, err := http.NewRequest("GET", url, nil)
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

func (r *Rabbit) CleanupRabbit() {
	r.Channel.Close()
	r.Connection.Close()
}
