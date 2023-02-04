package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

func createAsynqTask(channel string, message string) *asynq.Task {
	payload, _ := json.Marshal(map[string]string{
		"Inside SCHEDULER": message,
	})
	return asynq.NewTask(channel, payload)
}

// HandleIncomingMessages method  
// Listen to messages from the RabbitMQ
func (t *TaskHandler) HandleIncomingMessages() {
	go func() {
		for message := range t.Messages {
			log.Printf(">> Received TS message: %s\n", message.Body)
			t.TaskProducer.Enqueue(createAsynqTask("newsblast", "test"))
		}
	}()
}

func (handler *TaskHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case "Newsletter":
		log.Println("Handling newsletters")
		jsonString, _ := json.Marshal(map[string]string{
			"Hey": "Brother",
		})
		handler.MailerQueue.PublishMessage(string(jsonString))
		return nil

	case "type2":
		// process type2
	case "typeN":
		// process typeN

	default:
		return fmt.Errorf("unexpected task type: %q", t.Type())
	}
	return nil
}
