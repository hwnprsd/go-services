package main

import (
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"
	"github.com/hwnprsd/shared_types"
	"github.com/streadway/amqp"
)

func createAsynqTask(channel string, message string) *asynq.Task {
	payload, _ := json.Marshal(map[string]string{
		"Inside SCHEDULER": message,
	})
	return asynq.NewTask(channel, payload)
}

// HandleIncomingMessages method  
// Listen to messages from the RabbitMQ
func (t *TaskHandler) HandleMqMessages(channel *amqp.Channel) {

	messages, err := channel.Consume(
		SCHEDULER_QUEUE_NAME,
		"",
		false, // auto-ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   // arguments
	)

	if err != nil {
		log.Println("Error consuming for queue", SCHEDULER_QUEUE_NAME)
		log.Fatal(err)
	}
	go func() {
		for payload := range messages {
			log.SetPrefix("SCHEDULER: ")

			baseMessage := shared_types.MessagingBase{}
			defer func() {
				if err := recover(); err != nil {
					log.Println("panic occurred:", err)
					payload.Reject(false)
				}
			}()
			err := json.Unmarshal(payload.Body, &baseMessage)
			if err != nil {
				log.Printf("Error parsing JSON message. Please check what the sender sent! QUEUE - %s", payload.Body)
				payload.Reject(false)
				return
			}
			switch baseMessage.WorkType {
			case shared_types.WORK_TYPE_SCHEDULE_EMAIL:
				message := shared_types.ScheduleEmailsMessage{}
				json.Unmarshal(payload.Body, &message)
				taskData, _ := json.Marshal(message)
				task := asynq.NewTask(TASK_TYPE_EMAIL_LIST, taskData)
				log.Println(message.ScheduledTime.String())
				info, err := t.TaskProducer.Enqueue(task, asynq.ProcessAt(message.ScheduledTime))
				if err != nil {
					log.Println("Error enqueuing task", err)
				}
				log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
				break
			}
			payload.Ack(false)
		}
	}()
}

// - Handle Scheduled Mail Task
// Email template ID
// Email parameters + keys
// Schedule Date
// Mailing List ID
// Task ID

// - Handle Instant Mail Task
// Email template ID
// Email parameters + keys
// Schedule Date
// Mailing List ID
// Task ID
// Just pass the work to the mailer worker - to handle datafetch + deliver
