package mailer

import (
	"encoding/json"
	"log"

	"flaq.club/workers/utils"
	"github.com/hwnprsd/shared_types"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

type MailHandler struct {
	ApiQueue *utils.Queue
	Db       *gorm.DB
}

func NewMailHadler(ApiQueue *utils.Queue, Db *gorm.DB) *MailHandler {
	return &MailHandler{ApiQueue, Db}
}

func (h *MailHandler) HandleMessages(payload *amqp.Delivery) {
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
		payload.Ack(false)
	}
	switch baseMessage.WorkType {
	case shared_types.WORK_TYPE_SEND_MAIL:
		message := shared_types.SendMailMessage{}
		json.Unmarshal(payload.Body, &message)
		h.SendSingleEmail(message)
		break
	case shared_types.WORK_TYPE_SCHEDULE_EMAIL:
		message := shared_types.ScheduleEmailsMessage{}
		json.Unmarshal(payload.Body, &message)
		h.SendEmailToList(message)
		break
	case shared_types.WORK_TYPE_ADD_USER_TO_LIST:
		message := shared_types.AddUserListMessage{}
		json.Unmarshal(payload.Body, &message)
		h.AddUserToList(message)
		break
	}
	payload.Ack(false)
}
