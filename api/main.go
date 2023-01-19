package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"flaq.club/api/app"
	"flaq.club/api/controllers"
	"flaq.club/api/database"
	"flaq.club/api/messaging"
	"github.com/hibiken/asynq"
	"github.com/streadway/amqp"
)

// A list of task types.
const (
	TypeEmailDelivery = "email:deliver"
	TypeImageResize   = "image:resize"
)

type EmailDeliveryPayload struct {
	UserEmail  string
	TemplateID string
}

type ImageGenPayload struct {
	UserName string
}

func NewEmailDeliveryTask(userEmail string, tmplID string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailDeliveryPayload{UserEmail: userEmail, TemplateID: tmplID})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

func NewImageGenTask(userName string) (*asynq.Task, error) {
	payload, err := json.Marshal(ImageGenPayload{UserName: userName})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeImageResize, payload), nil
}

func main() {
	// With the instance and declare Queues that we can
	// publish and subscribe to.

	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	// Create a new RabbitMQ connection.
	rmqConnection, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer rmqConnection.Close()

	mq := messaging.Messaging{
		Connection: rmqConnection,
	}

	if err = mq.Setup(); err != nil {
		panic(err)
	}

	defer mq.MailerQueue.Channel.Close()

	fiberApp := controllers.New()

	db := database.Connect()

	app := app.New(&mq, fiberApp, db)

	// Setup the controller to have all app properties
	controller := controllers.Controller{
		app,
	}
	controller.SetupRoutes()

	log.Fatal(fiberApp.Listen(":3000"))

	// client := asynq.NewClient(asynq.RedisClientOpt{Addr: "127.0.0.1:6379"})
	// task, err := NewEmailDeliveryTask("ashwin@onpar.in", "some:template:id")
	// if err != nil {
	// 	log.Fatalf("could not create task: %v", err)
	// }
	// info, err := client.Enqueue(task)
	// if err != nil {
	// 	log.Fatalf("could not enqueue task: %v", err)
	// }
	// log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	//
	// task2, err2 := NewImageGenTask("ashwin@onpar.in")
	// if err2 != nil {
	// 	log.Fatalf("could not create task: %v", err)
	// }
	// info2, err := client.Enqueue(task2)
	// if err != nil {
	// 	log.Fatalf("could not enqueue task: %v", err)
	// }
	// log.Printf("enqueued task: id=%s queue=%s", info2.ID, info2.Queue)
	//
	for range time.Tick(time.Second * 10) {

	}

}
