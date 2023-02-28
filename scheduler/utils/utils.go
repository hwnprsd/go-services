package utils

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type Queue struct {
	Name    string
	Channel *amqp.Channel
}

func (q *Queue) PublishMessage(payloadMap interface{}) error {

	jsonString, _ := json.Marshal(payloadMap)

	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(string(jsonString)),
	}
	if err := q.Channel.Publish(
		"",     // exchange
		q.Name, // queue name
		false,  // mandatory
		false,  // immediate
		message,
	); err != nil {
		return err
	}
	return nil
}

func CreateQueue(conn amqp.Connection, queueName string) (*Queue, func()) {
	// Open a channel
	channel, err := conn.Channel()
	if err != nil {
		log.Println("Error creating a channel while creating queue ", queueName)
		panic(err)
	}
	_, err = channel.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // arguments
	)
	if err != nil {
		log.Println("Error a declaring queue: ", queueName)
		panic(err)
	}
	return &Queue{
			Name:    queueName,
			Channel: channel,
		}, func() {
			channel.Close()
		}
}
