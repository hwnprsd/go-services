package shared_types

const WORK_TYPE_SEND_MAIL = "SEND_MAIL"

type SendMailMessage struct {
	MessagingBase
	EmailAddress   string
	Subject        string
	BodyTemplateID uint
	TemplateValues map[string]string
}

func NewSendMailMessage(email string, subject string, templateId uint, values map[string]string) *SendMailMessage {
	return &SendMailMessage{
		MessagingBase: MessagingBase{
			WorkType: WORK_TYPE_SEND_MAIL,
		},
		EmailAddress:   email,
		Subject:        subject,
		BodyTemplateID: templateId,
		TemplateValues: values,
	}
}
