package mailer

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"flaq.club/workers/models"
	"github.com/hwnprsd/shared_types"
	"gorm.io/gorm"
)

func SendSingleEmail(db *gorm.DB, data shared_types.SendMailMessage) {
	log.Println("Attempting to send single email")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USER")

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	emailTemplate := models.EmailTemplate{}
	result := db.Model(&emailTemplate).Where("id = ?", data.BodyTemplateID).First(&emailTemplate)
	if result.Error != nil {
		log.Println("Error occoured fetching email template")
		log.Println(result.Error)
		return
	}

	log.Println(data.TemplateValues...)
	log.Println(data.TemplateValues[len(data.TemplateValues)-1])

	message := fmt.Sprintf(emailTemplate.TemplateString, data.TemplateValues...)

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"

	emailBody :=
		"To: " + data.EmailAddress + "\r\n" +
			"From: Flaq<welcome@flaq.club>\r\n" +
			"Subject: " + data.Subject + "\r\n" +
			mime +
			"\r\n" +
			message

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, "welcome@flaq.club", []string{data.EmailAddress}, []byte(emailBody))

	if err != nil {
		log.Println("Error occoured sending email")
		log.Println(err)
		return
	}
	log.Println("Email sent succesfully")

}
