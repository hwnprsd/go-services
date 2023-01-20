package messaging

import (
	"github.com/streadway/amqp"
)

type Messaging struct {
	Connection *amqp.Connection
	// Declare all the queues required for the application
	MailerQueue *Queue
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
		"",       // exchange
		"Mailer", // queue name
		false,    // mandatory
		false,    // immediate
		message,
	); err != nil {
		panic(err)
	}
}

// Setup method  
func (m *Messaging) Setup() error {
	queue, err := m.CreateQueue("Mailer")
	if err != nil {
		return err
	}
	m.MailerQueue = queue
	return nil
}

// CreateQueue method  
func (m *Messaging) CreateQueue(queueName string) (*Queue, error) {
	// Open a channel
	channel, err := m.Connection.Channel()
	if err != nil {
		return nil, err
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
		return nil, err
	}
	return &Queue{
		Name:    queueName,
		Channel: *channel,
	}, nil
}
