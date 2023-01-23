package main

import (
	"log"
	"os"

	"flaq.club/workers/pkgs/nft"
	"github.com/streadway/amqp"
)

const QUEUE_NAME_MAILER = "mailer"
const QUEUE_NAME_NFT = "nft"

// main function  
func main() {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	connectRabbitMq, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMq.Close()
	log.Println("Mailer Listener Started")
	queues := []string{QUEUE_NAME_MAILER, QUEUE_NAME_NFT}
	forever := make(chan bool)
	for _, queue := range queues {
		channel, err := connectRabbitMq.Channel()
		if err != nil {
			panic(err)
		}
		defer channel.Close()
		defer log.Printf("QUEUE CLOSED! QUEUE CLOSED! QUEUE CLOSED! QUEUE CLOSED!")
		go ProcessQueue(*channel, queue)
	}
	<-forever
}

func ProcessQueue(channel amqp.Channel, queueName string) {

	channel.QueueDeclare(
		queueName,
		true,  // durable
		false, // auto delete
		false, // exclusive
		false, // no wait
		nil,   // arguments
	)

	// Subscribing to QueueService1 for getting messages.
	messages, err := channel.Consume(
		queueName, // queue name
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       // arguments
	)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
	go func() {
		for message := range messages {
			if queueName == QUEUE_NAME_NFT {
				nft.HandleMessages(&message)
			}
			if queueName == QUEUE_NAME_MAILER {

			}

		}
	}()
}
