package main

import (
	"log"
	"os"

	"flaq.club/workers/database"
	"flaq.club/workers/pkgs/mailer"
	"flaq.club/workers/pkgs/nft"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
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
	db := database.Connect()
	for _, queue := range queues {
		channel, err := connectRabbitMq.Channel()
		if err != nil {
			panic(err)
		}
		defer channel.Close()
		defer log.Printf("QUEUE CLOSED! QUEUE CLOSED! QUEUE CLOSED! QUEUE CLOSED!")
		go ProcessQueue(*channel, queue, db)
	}
	<-forever
}

func ProcessQueue(channel amqp.Channel, queueName string, db *gorm.DB) {

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

	nftMintHander := nft.NewNftMintHandler(db)

	go func() {
		for message := range messages {
			if queueName == QUEUE_NAME_NFT {
				log.Println("Handling NFT Message")
				nftMintHander.HandleMessages(&message)
			}
			if queueName == QUEUE_NAME_MAILER {
				log.Println("Handling Mailer Message")
				mailer.HandleMessages(&message, db)
			}

		}
	}()
}
