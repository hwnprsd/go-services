package mailer

import (
	"encoding/json"
	"log"

	"github.com/hwnprsd/shared_types"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

func HandleMessages(payload *amqp.Delivery, db *gorm.DB) {
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
	}
	switch baseMessage.WorkType {
	case shared_types.WORK_TYPE_SEND_MAIL:
		message := shared_types.SendMailMessage{}
		json.Unmarshal(payload.Body, &message)
		SendSingleEmail(db, message)
	}
	payload.Ack(false)
}
