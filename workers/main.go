package main

import (
	"log"
	"os"

	"flaq.club/workers/database"
	"flaq.club/workers/pkgs/gif"
	"flaq.club/workers/pkgs/mailer"
	"flaq.club/workers/pkgs/nft"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

const QUEUE_NAME_MAILER = "mailer"
const QUEUE_NAME_NFT = "nft"
const QUEUE_NAME_GIF = "gif"

// main function  
func main() {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	connectRabbitMq, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMq.Close()
	log.Println("Mailer Listener Started")
	queues := []string{QUEUE_NAME_MAILER, QUEUE_NAME_NFT, QUEUE_NAME_GIF}
	forever := make(chan bool)
	db := database.Connect()
	for _, queue := range queues {
		channel, err := connectRabbitMq.Channel()
		if err != nil {
			panic(err)
		}
		defer channel.Close()
		defer log.Printf("QUEUE CLOSED! QUEUE CLOSED! QUEUE CLOSED! QUEUE CLOSED!")
		go ProcessQueue(connectRabbitMq, *channel, queue, db)
	}
	<-forever
}

func ProcessQueue(connection *amqp.Connection, channel amqp.Channel, queueName string, db *gorm.DB) {

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
				mailerChannel, err := connection.Channel()
				defer mailerChannel.Close()
				if err != nil {
					log.Fatal("Error creating a channel")
				}
				mailerChannel.QueueDeclare(
					QUEUE_NAME_MAILER,
					true,  // durable
					false, // auto delete
					false, // exclusive
					false, // no wait
					nil,   // arguments
				)
				nftMintHander := nft.NewNftMintHandler(mailerChannel, db)
				log.Printf("Handling NFT Message from := %s", message.Timestamp)
				nftMintHander.HandleMessages(&message)
			}
			if queueName == QUEUE_NAME_MAILER {
				log.Printf("Handling Mailer Message from := %s", message.Timestamp)
				mailer.HandleMessages(&message, db)
			}
			if queueName == QUEUE_NAME_GIF {
				nftChannel, err := connection.Channel()
				defer nftChannel.Close()
				if err != nil {
					log.Fatal("Error creating a channel")
				}
				nftChannel.QueueDeclare(
					QUEUE_NAME_NFT,
					true,  // durable
					false, // auto delete
					false, // exclusive
					false, // no wait
					nil,   // arguments
				)
				gifHandler := gif.NewCreateGifHandler(nftChannel, db)
				log.Printf("Handling GIf Message from := %s", message.Timestamp)
				gifHandler.HandleMessages(&message)
			}
		}
	}()
}
