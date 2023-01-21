package main

import (
	"context"
	"log"
	"os"

	"github.com/hibiken/asynq"
	"github.com/streadway/amqp"
)

const SCHEDULER_QUEUE_NAME = "scheduler"

type Queue struct {
	Name    string
	Channel amqp.Channel
}

func (q *Queue) PublishMessage(payload string) {
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(payload),
	}
	if err := q.Channel.Publish(
		"",     // exchange
		q.Name, // queue name
		false,  // mandatory
		false,  // immediate
		message,
	); err != nil {
		panic(err)
	}
}

type TaskHandler struct {
	Connection   *amqp.Connection
	MailerQueue  *Queue
	ApiQueue     *Queue
	Messages     <-chan amqp.Delivery
	TaskProducer asynq.Client
}

func myHandler(ctx context.Context, t *asynq.Task) error {
	log.Println("NEW MESSAGE")
	log.Println(t.Payload())
	return nil
}

func (t *TaskHandler) createConsumer() func() {
	channel, err := t.Connection.Channel()
	if err != nil {
		panic(err)
	}
	channel.QueueDeclare(
		SCHEDULER_QUEUE_NAME,
		true,  // durable
		false, // auto delete
		false, // exclusive
		false, // no wait
		nil,   // arguments
	)
	messages, err := channel.Consume(
		SCHEDULER_QUEUE_NAME,
		"",
		true,  // auto-ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   // arguments
	)
	t.Messages = messages
	return func() {
		channel.Close()
	}
}

func SetupHandler(conn *amqp.Connection) (*TaskHandler, func()) {
	channel1, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	channel2, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	taskClient := *asynq.NewClient(asynq.RedisClientOpt{Addr: os.Getenv("REDIS_URL")})

	tH := TaskHandler{
		Connection: conn,
		MailerQueue: &Queue{
			Channel: *channel1,
			Name:    "mailer",
		},
		ApiQueue: &Queue{
			Channel: *channel2,
			Name:    "api",
		},
		TaskProducer: taskClient,
	}
	closeFunc := tH.createConsumer()
	return &tH, func() {
		closeFunc()
		channel1.Close()
		channel2.Close()
		taskClient.Close()
	}

}

func main() {
	redisAddr := os.Getenv("REDIS_URL")

	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	rmqConnection, err := amqp.Dial(amqpServerURL)
	defer rmqConnection.Close()

	taskHandler, closeFunc := SetupHandler(rmqConnection)
	defer closeFunc()

	if err != nil {
		panic(err)
	}

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	log.Printf("Running asynq server")
	taskHandler.HandleIncomingMessages()
	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc("newsblast", myHandler)
	// mux.Handle(tasks.TypeImageResize, tasks.NewImageProcessor())
	// ...register other handlers...

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
