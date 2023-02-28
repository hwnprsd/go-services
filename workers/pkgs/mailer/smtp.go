package mailer

import (
	"encoding/json"
	"log"
	"net/smtp"
	"strings"
	"sync"

	"flaq.club/workers/models"
	"flaq.club/workers/utils"
	"github.com/hwnprsd/shared_types"
	"github.com/hwnprsd/shared_types/status"
	"gorm.io/datatypes"
)

func (h *MailHandler) SendSingleEmail(message shared_types.SendMailMessage) {
	log.SetPrefix("MAIL_HANDLER: ")
	log.Println("Attempting to send single email")

	startMessage := shared_types.NewApiCallbackMessage(message.TaskId, status.POAP_MAILER_START, "")
	h.ApiQueue.PublishMessage(startMessage)

	emailTemplate := models.EmailTemplate{}
	emailMessage := emailTemplate.TemplateString
	result := h.Db.Model(&emailTemplate).Where("id = ?", message.BodyTemplateID).First(&emailTemplate)
	if result.Error != nil {
		log.Println("Error occoured fetching email template")
		log.Println(result.Error)
		return
	}

	err := h.sendMail(emailMessage, message.EmailAddress, emailTemplate.Subject, message.TemplateValues, &sync.WaitGroup{})
	if err != nil {
		log.Println("Error sending email")
	}

	endMessage := shared_types.NewApiCallbackMessage(message.TaskId, status.POAP_MAILER_COMPLETE, "")
	h.ApiQueue.PublishMessage(endMessage)
	log.Println("Email sent succesfully")

}

func (h *MailHandler) SendEmailToList(message shared_types.ScheduleEmailsMessage) {
	log.SetPrefix("MULTI_MAIL_HANDLER: ")
	campaign := models.MailingCampaign{}
	res := h.Db.Preload("EmailTemplate").Preload("MailingList.Users").Where("id = ?", message.CampaignId).First(&campaign)
	if res.Error != nil {
		log.Println("Error finding campaign for id", message.CampaignId)
		log.Println("Failing Silently")
		log.Println("-----------------------------")
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(len(campaign.MailingList.Users))

	for _, user := range campaign.MailingList.Users {
		log.Println("Current User", user.EmailAddress)
		// Check if the user's meta value matches the campaign's requirement
		if campaign.MetaKey == "" {
			// Ignore the meta
			go h.sendMail(campaign.EmailTemplate.TemplateString, user.EmailAddress, campaign.EmailTemplate.Subject, message.TemplateValues, &wg)
		} else {
			bytes, err := user.MetaData.MarshalJSON()
			if err != nil {
				log.Println("Error unmarshalling user json", message.CampaignId)
				log.Println("Failing Silently")
				log.Println("-----------------------------")
				wg.Done()
				continue
			}
			var userMetaData map[string]uint
			json.Unmarshal(bytes, &userMetaData)
			log.Println("UserMeta - ", userMetaData)
			log.Println(user.MetaData.String())
			if userMetaData[campaign.MetaKey] != campaign.MetaValue {
				// Need not init if the meta doesn't match
				if !campaign.ShouldInitMeta {
					wg.Done()
					continue
				} else if campaign.ShouldInitMeta && userMetaData[campaign.MetaKey] == 0 {
					if userMetaData == nil {
						userMetaData = make(map[string]uint)
					}
					userMetaData[campaign.MetaKey] = campaign.MetaValue
					bytes, _ := json.Marshal(userMetaData)
					res := h.Db.Model(&user).Where("email_address = ?", user.EmailAddress).Update("meta_data", datatypes.JSON(bytes))

					if res.Error != nil {
						log.Println("Error updating user's initital metadata")
					} else {
					}
				} else {
					// Already initied, invalid user
					log.Println("User's Meta already inited, and is invalid")
					wg.Done()
					continue

				}
			}
			go h.sendMail(campaign.EmailTemplate.TemplateString, user.EmailAddress, campaign.EmailTemplate.Subject, message.TemplateValues, &wg)
			if campaign.ShouldIncrementMeta {
				if userMetaData == nil {
					userMetaData = make(map[string]uint)
				}
				userMetaData[campaign.MetaKey] = campaign.MetaValue + 1
				bytes, _ := json.Marshal(userMetaData)
				res := h.Db.Model(&user).Where("email_address = ?", user.EmailAddress).Update("meta_data", datatypes.JSON(bytes))
				if res.Error != nil {
					log.Println("Error updating user's metadata")
				} else {
					log.Println("User's metadata updated")

				}
			}
		}
	}
	wg.Wait()
	log.Println("Wait group released")
}

func (h *MailHandler) sendMail(emailMessage, emailAddress, subject string, templateValues map[string]string, wg *sync.WaitGroup) error {
	defer wg.Done()
	log.Println("Sending email to ", emailAddress)
	smtpHost, smtpUsername, smtpPassword, smtpPort := utils.GetMailerEnv()
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
	for key, value := range templateValues {
		emailMessage = strings.ReplaceAll(emailMessage, "{{ "+key+" }}", value)
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"

	emailBody :=
		"To: " + emailAddress + "\r\n" +
			"From: Flaq<welcome@flaq.club>\r\n" +
			"Subject: " + subject + "\r\n" +
			mime +
			"\r\n" +
			emailMessage

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, "welcome@flaq.club", []string{emailAddress}, []byte(emailBody))

	if err != nil {
		log.Println("Error occoured sending email")
		log.Println("SMTP HOST", smtpHost)
		return err
	}
	return nil
}
