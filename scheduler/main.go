package main

import (
	"context"
	"log"
	"os"

	"flaq.club/scheduler/utils"
	"github.com/hibiken/asynq"
	"github.com/streadway/amqp"
)

const SCHEDULER_QUEUE_NAME = "scheduler"
const TASK_TYPE_EMAIL_LIST = "EMAIL_LIST"

type Queue struct {
	Name    string
	Channel *amqp.Channel
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
	MqConnection *amqp.Connection
	MailerQueue  *utils.Queue
	ApiQueue     *utils.Queue
	TaskProducer *asynq.Client
}

func (th *TaskHandler) handleEmailListTask(ctx context.Context, t *asynq.Task) error {
	log.SetPrefix("EMAIL_LIST_HANDLER:")
	log.Println("Handling Email Task")
	log.Println(t.Payload())
	return nil
}

func (t *TaskHandler) createConsumer(queueName string) (*amqp.Channel, func()) {
	channel, err := t.MqConnection.Channel()
	if err != nil {
		panic(err)
	}
	channel.QueueDeclare(
		queueName,
		true,  // durable
		false, // auto delete
		false, // exclusive
		false, // no wait
		nil,   // arguments
	)
	return channel, func() {
		channel.Close()
	}
}

// Setup the messaging queu, opening up the different channels, etc
func SetupMqHandler(conn *amqp.Connection) (*TaskHandler, func()) {
	mailerQueue, cancelFunc1 := utils.CreateQueue(*conn, "mailer")
	apiQueue, cancelFunc2 := utils.CreateQueue(*conn, "api")

	taskClient := asynq.NewClient(asynq.RedisClientOpt{Addr: os.Getenv("REDIS_URL")})

	tH := TaskHandler{
		MqConnection: conn,
		MailerQueue:  mailerQueue,
		ApiQueue:     apiQueue,
		TaskProducer: taskClient,
	}
	return &tH, func() {
		cancelFunc1()
		cancelFunc2()
		taskClient.Close()
	}

}

func main() {
	redisAddr := os.Getenv("REDIS_URL")

	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	rmqConnection, err := amqp.Dial(amqpServerURL)
	defer rmqConnection.Close()

	taskHandler, closeFunc := SetupMqHandler(rmqConnection)
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

	go func() {

		forever := make(chan bool)
		channel, channelCloseFunc := taskHandler.createConsumer(SCHEDULER_QUEUE_NAME)
		defer channelCloseFunc()
		go taskHandler.HandleMqMessages(channel)
		<-forever
	}()
	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(TASK_TYPE_EMAIL_LIST, taskHandler.handleEmailListTask)
	// mux.Handle(tasks.TypeImageResize, tasks.NewImageProcessor())
	// ...register other handlers...

	log.Printf("Running asynq server")
	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
	log.Println("Closing async service")
}
