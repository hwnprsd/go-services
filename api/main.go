package main

import (
	"log"
	"os"

	"flaq.club/api/app"
	"flaq.club/api/controllers"
	"flaq.club/api/database"
	"flaq.club/api/messaging"
	"github.com/streadway/amqp"
)

type EmailDeliveryPayload struct {
	UserEmail  string
	TemplateID string
}

type ImageGenPayload struct {
	UserName string
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

	closeFunc := mq.Setup()

	defer closeFunc()

	fiberApp := controllers.New()

	db := database.Connect()

	app := app.New(&mq, fiberApp, db)

	// Setup the controller to have all app properties
	controller := controllers.Controller{
		app,
	}
	controller.SetupRoutes()

	log.Fatal(fiberApp.Listen(":3000"))

}
