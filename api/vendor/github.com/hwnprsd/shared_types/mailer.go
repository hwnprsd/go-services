package shared_types

const (
	WORK_TYPE_SEND_MAIL           = "SEND_MAIL"
	WORK_TYPE_ADD_USER_TO_LIST    = "MAILER_ADD_USER_LIST"
	WORK_TYPE_SCHEDULE_CAMPAIGN   = "MAILER_SCHEDULE_CAMPAIGN"
	WORK_TYPE_REMOVE_USER_TO_LIST = "MAILER_REMOVE_USER_LIST"
	WORK_TYPE_TRIGGER_CAMPAIGN    = "MAILER_TRIGGER_CAMPAIGN"
	WORK_TYPE_CAMPAIGN_SUCCESS    = "MAILER_CAMPAIGN_SUCCESS"
)

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

type AddUserListMessage struct {
	MessagingBase
	EmailAddress string
	Name         string
	ListName     string
}

type RemoveUserListMessage struct {
	MessagingBase
	EmailAddress string
	Name         string
	ListName     string
}

type CreateCampaignMessage struct {
	MessagingBase
	ListName string
}
