package main

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

const QUEUE_NAME = "mailer"

// main function  
func main() {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	connectRabbitMq, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMq.Close()
	channel, err := connectRabbitMq.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	channel.QueueDeclare(
		QUEUE_NAME,
		true,  // durable
		false, // auto delete
		false, // exclusive
		false, // no wait
		nil,   // arguments
	)

	// Subscribing to QueueService1 for getting messages.
	messages, err := channel.Consume(
		QUEUE_NAME, // queue name
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // arguments
	)
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		panic(err)
	}

	fmt.Println("Mailer Listener Started")
	forever := make(chan bool)
	go func() {
		for message := range messages {
			log.Printf(">> Received message: %s\n", message.Body)
		}
	}()
	<-forever
}
