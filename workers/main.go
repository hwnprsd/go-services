package main

import (
	"log"
	"os"

	"flaq.club/workers/database"
	"flaq.club/workers/pkgs/gif"
	"flaq.club/workers/pkgs/mailer"
	"flaq.club/workers/pkgs/nft"
	"flaq.club/workers/utils"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

const QUEUE_NAME_MAILER = "mailer"
const QUEUE_NAME_NFT = "nft"
const QUEUE_NAME_GIF = "gif"
const QUEUE_NAME_API = "api"

// main function  
func main() {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	connectRabbitMq, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMq.Close()
	log.Println("Mailer Listener Started")
	forever := make(chan bool)
	db := database.Connect()
	apiQueue, cancelFunc1 := utils.CreateQueue(*connectRabbitMq, QUEUE_NAME_API)
	nftQueue, cancelFunc2 := utils.CreateQueue(*connectRabbitMq, QUEUE_NAME_NFT)
	gifQueue, cancelFunc3 := utils.CreateQueue(*connectRabbitMq, QUEUE_NAME_GIF)
	mailerQueue, cancelFunc4 := utils.CreateQueue(*connectRabbitMq, QUEUE_NAME_MAILER)
	defer func() {
		cancelFunc1()
		cancelFunc2()
		cancelFunc3()
		cancelFunc4()
		log.Println("Cancelling all functions")
	}()

	handler := WorkHandler{
		ApiQueue:    apiQueue,
		NFTQueue:    nftQueue,
		MailerQueue: mailerQueue,
		GifQueue:    gifQueue,
		Db:          db,
	}
	go handler.ProcessQueue(handler.GifQueue)
	go handler.ProcessQueue(handler.NFTQueue)
	go handler.ProcessQueue(handler.MailerQueue)
	<-forever
}

type WorkHandler struct {
	ApiQueue    *utils.Queue
	NFTQueue    *utils.Queue
	MailerQueue *utils.Queue
	GifQueue    *utils.Queue
	Db          *gorm.DB
}

func (h *WorkHandler) ProcessQueue(queue *utils.Queue) {
	// Subscribing to QueueService1 for getting messages.
	log.Println("Processing queue for", queue.Name)
	messages, err := queue.Channel.Consume(
		queue.Name, // queue name
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // arguments
	)
	if err != nil {
		log.Println("Error consuming for queue", queue.Name)
		log.Fatal(err)
	}
	go func() {
		// Create these channels to help publish messages ONLY
		for message := range messages {
			if queue.Name == QUEUE_NAME_NFT {
				nftMintHander := nft.NewNftMintHandler(h.ApiQueue, h.MailerQueue, h.Db)
				log.Printf("Handling NFT Message from := %s", message.Timestamp)
				nftMintHander.HandleMessages(&message)
			}
			if queue.Name == QUEUE_NAME_MAILER {
				log.Printf("Handling Mailer Message from := %s", message.Timestamp)
				mailHandler := mailer.NewMailHadler(h.ApiQueue, h.Db)
				mailHandler.HandleMessages(&message)
			}
			if queue.Name == QUEUE_NAME_GIF {
				gifHandler := gif.NewCreateGifHandler(h.ApiQueue, h.NFTQueue, h.Db)
				log.Printf("Handling GIf Message from := %s", message.Timestamp)
				gifHandler.HandleMessages(&message)
			}
		}
		log.Println("Closing queue handler", queue.Name)
		log.Println("Restarting queue")
		h.ProcessQueue(queue)
	}()
}
