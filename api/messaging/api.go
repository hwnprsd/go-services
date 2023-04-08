package messaging

import (
	"encoding/json"
	"fmt"
	"log"

	"flaq.club/api/models"
	"github.com/hwnprsd/shared_types"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

func HandleApiMessages(messages <-chan amqp.Delivery, db *gorm.DB) {
	go func() {
		for message := range messages {
			log.Println("Handling messages coming into the API Queue", message.Timestamp)
			// Handle messages
			baseMessage := shared_types.MessagingBase{}
			defer func() {
				if err := recover(); err != nil {
					log.Println("panic occurred:", err)
					message.Reject(false)
				}
			}()
			fmt.Println(string(message.Body))
			err := json.Unmarshal(message.Body, &baseMessage)
			if err != nil {
				log.Printf("Error parsing JSON message. Please check what the sender sent! QUEUE - %s\n", "_")
				message.Reject(false)
				return
			}
			switch baseMessage.WorkType {
			case shared_types.WORK_TYPE_API_CALLBACK:
				apiCallbackMessage := shared_types.ApiCallback{}
				json.Unmarshal(message.Body, &apiCallbackMessage)

				task := models.Task{}
				log.Println("Task ID", apiCallbackMessage.TaskID)
				log.Println("Status", apiCallbackMessage.Status)
				log.Println("Data", apiCallbackMessage.Data)
				res := db.Model(&task).Where("id = ?", apiCallbackMessage.TaskID).Updates(map[string]interface{}{"status": apiCallbackMessage.Status, "data": apiCallbackMessage.Data})
				// // TODO Update ID as well
				if res.Error != nil {
					log.Println("Invalid Task ID")
					log.Println(res.Error)
				}
				message.Ack(false)
				break

			}
		}

	}()
}
