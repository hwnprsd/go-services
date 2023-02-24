package utils

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

func PublishMessage(channel *amqp.Channel, name string, payloadMap interface{}) error {

	jsonString, _ := json.Marshal(payloadMap)

	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(string(jsonString)),
	}
	if err := channel.Publish(
		"",    // exchange
		name,  // queue name
		false, // mandatory
		false, // immediate
		message,
	); err != nil {
		return err
	}
	return nil
}
