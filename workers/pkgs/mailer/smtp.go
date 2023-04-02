package mailer

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync"

	"flaq.club/workers/models"
	"flaq.club/workers/utils"
	"github.com/go-mail/mail"
	"github.com/hwnprsd/shared_types"
	"github.com/hwnprsd/shared_types/status"
	"gorm.io/datatypes"
	"gorm.io/gorm/clause"
)

func (h *MailHandler) AddUserToList(message shared_types.AddUserListMessage) {
	mailingList := models.MailingList{}
	if res := h.Db.Where("list_name = ?", message.ListName).First(&mailingList); res.Error != nil {
		log.Println("Error finding list with name", message.ListName)
		return
	}
	user := models.MailingUser{}
	log.Println("Trying to add -", message.EmailAddress)
	count := int64(0)
	res := h.Db.Model(&user).Where("email_address = ?", message.EmailAddress).Count(&count)
	if res.Error != nil {
		log.Println("Error Fetching User Count - Email =", message.EmailAddress, res.Error)
		return
	}
	if count == 0 {
		// Create user
		user.EmailAddress = message.EmailAddress
		user.Name = message.Name
		h.Db.Clauses(clause.Returning{}).Create(&user)
	}
	if err := h.Db.Model(&user).Association("Lists").Append(&mailingList); err != nil {
		log.Println("Error adding user to a list")
		return
	}

}

func (h *MailHandler) SendSingleEmail(message shared_types.SendMailMessage) {
	log.SetPrefix("MAIL_HANDLER: ")
	log.Println("Attempting to send single email")

	startMessage := shared_types.NewApiCallbackMessage(message.TaskId, status.POAP_MAILER_START, "")
	h.ApiQueue.PublishMessage(startMessage)

	emailTemplate := models.EmailTemplate{}
	result := h.Db.Model(&emailTemplate).Where("id = ?", message.BodyTemplateID).First(&emailTemplate)
	if result.Error != nil {
		log.Println("Error occoured fetching email template")
		log.Println(result.Error)
		return
	}
	emailMessage := emailTemplate.TemplateString
	wg := &sync.WaitGroup{}
	wg.Add(1)
	err := h.sendMail(emailMessage, message.EmailAddress, emailTemplate.Subject, message.TemplateValues, wg)
	if err != nil {
		log.Println("Error sending email - ", err)
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
	smtpHost, smtpUsername, smtpPassword, smtpPort := utils.GetMailerEnv()
	defer wg.Done()
	m := mail.NewMessage()
	m.SetHeader("From", "Flaq<welcome@flaq.club>")
	m.SetHeader("To", emailAddress)
	m.SetHeader("Subject", subject)

	for key, value := range templateValues {
		emailMessage = strings.ReplaceAll(emailMessage, "{{ "+key+" }}", value)
	}

	m.SetBody("text/html", emailMessage)

	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		log.Println("Error converting port to int inside mailer")
		return err
	}

	d := mail.NewDialer(smtpHost, port, smtpUsername, smtpPassword)
	if err := d.DialAndSend(m); err != nil {
		log.Println("Error occured sending email")
		log.Println(err)
		return err
	}

	return nil
}
