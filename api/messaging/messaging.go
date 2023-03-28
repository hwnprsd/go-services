package messaging

import (
	"encoding/json"

	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

type Messaging struct {
	Connection *amqp.Connection
	// Declare all the queues required for the application
	MailerQueue    *Queue
	SchedulerQueue *Queue
	NftQueue       *Queue
	GifQueue       *Queue
	ApiQueue       *Queue
	ScraperQueue   *Queue
	GPTQueue       *Queue
}

type Queue struct {
	Name    string
	Channel amqp.Channel
}

func (q *Queue) PublishMessage(payloadMap interface{}) error {

	jsonString, _ := json.Marshal(payloadMap)

	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(string(jsonString)),
	}
	if err := q.Channel.Publish(
		"",     // exchange
		q.Name, // queue name
		false,  // mandatory
		false,  // immediate
		message,
	); err != nil {
		return err
	}
	return nil
}

// Setup method  
func (m *Messaging) Setup() func() {
	queue1, closeFunc1 := m.CreateQueue("mailer")
	queue2, closeFunc2 := m.CreateQueue("scheduler")
	queue3, closeFunc3 := m.CreateQueue("nft")
	queue4, closeFunc4 := m.CreateQueue("gif")
	queue5, closeFunc5 := m.CreateQueue("api")
	queue6, closeFunc6 := m.CreateQueue("scraper")
	queue7, closeFunc7 := m.CreateQueue("gptx")
	m.MailerQueue = queue1
	m.SchedulerQueue = queue2
	m.NftQueue = queue3
	m.GifQueue = queue4
	m.ApiQueue = queue5
	m.ScraperQueue = queue6
	m.GPTQueue = queue7

	return func() {
		closeFunc1()
		closeFunc2()
		closeFunc3()
		closeFunc4()
		closeFunc5()
		closeFunc6()
		closeFunc7()
	}
}

func (m *Messaging) SetupApiMessageListener(db *gorm.DB) {
	messages, err := m.ApiQueue.Channel.Consume(
		"api",
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   // arguments
	)
	if err != nil {
		panic(err)
	}

	HandleApiMessages(messages, db)
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
