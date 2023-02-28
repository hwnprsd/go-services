package mailer

import (
	"log"
	"net/smtp"
	"os"
	"strings"

	"flaq.club/workers/models"
	"github.com/hwnprsd/shared_types"
	"github.com/hwnprsd/shared_types/status"
)

func (h *MailHandler) SendSingleEmail(message shared_types.SendMailMessage) {
	log.SetPrefix("MAIL_HANDLER: ")
	log.Println("Attempting to send single email")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USER")

	startMessage := shared_types.NewApiCallbackMessage(message.TaskId, status.POAP_MAILER_START, "")
	h.ApiQueue.PublishMessage(startMessage)

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	emailTemplate := models.EmailTemplate{}
	result := h.Db.Model(&emailTemplate).Where("id = ?", message.BodyTemplateID).First(&emailTemplate)
	if result.Error != nil {
		log.Println("Error occoured fetching email template")
		log.Println(result.Error)
		return
	}

	emailMessage := emailTemplate.TemplateString
	for key, value := range message.TemplateValues {
		emailMessage = strings.ReplaceAll(emailMessage, "{{ "+key+" }}", value)
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"

	emailBody :=
		"To: " + message.EmailAddress + "\r\n" +
			"From: Flaq<welcome@flaq.club>\r\n" +
			"Subject: " + message.Subject + "\r\n" +
			mime +
			"\r\n" +
			emailMessage

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, "welcome@flaq.club", []string{message.EmailAddress}, []byte(emailBody))

	if err != nil {
		log.Println("Error occoured sending email")
		log.Println("SMTP HOST", smtpHost)
		panic(err)
	}
	endMessage := shared_types.NewApiCallbackMessage(message.TaskId, status.POAP_MAILER_COMPLETE, "")
	h.ApiQueue.PublishMessage(endMessage)
	log.Println("Email sent succesfully")

}
