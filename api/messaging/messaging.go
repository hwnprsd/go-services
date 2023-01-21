package messaging

import (
	"github.com/streadway/amqp"
)

type Messaging struct {
	Connection *amqp.Connection
	// Declare all the queues required for the application
	MailerQueue    *Queue
	SchedulerQueue *Queue
}

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

// Setup method  
func (m *Messaging) Setup() func() {
	queue1, closeFunc1 := m.CreateQueue("mailer")
	queue2, closeFunc2 := m.CreateQueue("scheduler")
	m.MailerQueue = queue1
	m.SchedulerQueue = queue2

	return func() {
		closeFunc1()
		closeFunc2()
	}
}

// CreateQueue method  
func (m *Messaging) CreateQueue(queueName string) (*Queue, func()) {
	// Open a channel
	channel, err := m.Connection.Channel()
	if err != nil {
		panic(err)
	}
	_, err = channel.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // arguments
	)
	if err != nil {
		panic(err)
	}
	return &Queue{
			Name:    queueName,
			Channel: *channel,
		}, func() {
			channel.Close()
		}
}