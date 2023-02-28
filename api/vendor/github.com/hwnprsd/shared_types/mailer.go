package shared_types

import "time"

const (
	WORK_TYPE_SEND_MAIL           = "SEND_MAIL"
	WORK_TYPE_ADD_USER_TO_LIST    = "MAILER_ADD_USER_LIST"
	WORK_TYPE_SCHEDULE_CAMPAIGN   = "MAILER_SCHEDULE_CAMPAIGN"
	WORK_TYPE_REMOVE_USER_TO_LIST = "MAILER_REMOVE_USER_LIST"
	WORK_TYPE_TRIGGER_CAMPAIGN    = "MAILER_TRIGGER_CAMPAIGN"
	WORK_TYPE_CAMPAIGN_SUCCESS    = "MAILER_CAMPAIGN_SUCCESS"

	WORK_TYPE_SCHEDULE_EMAIL = "SCHEDULE_EMAIL"
	WORK_TYPE_SEND_MAIL_LIST = "SEND_MAIL_LIST"
)

type SendMailMessage struct {
	MessagingBase
	EmailAddress   string
	Subject        string
	BodyTemplateID uint
	TemplateValues map[string]string
}

func NewSendMailMessage(taskId uint, email string, subject string, templateId uint, values map[string]string) *SendMailMessage {
	return &SendMailMessage{
		*NewMessagingBase(taskId, WORK_TYPE_SEND_MAIL),
		email,
		subject, // Subject should come from the DB not via a message
		templateId,
		values,
	}
}

type ScheduleEmailsMessage struct {
	MessagingBase
	CampaignId     uint
	ScheduledTime  time.Time
	TemplateValues map[string]string
}

func NewScheduleEmailMessage(taskId, campaignId uint, scheduledTime time.Time, templateValues map[string]string) *ScheduleEmailsMessage {
	return &ScheduleEmailsMessage{
		*NewMessagingBase(taskId, WORK_TYPE_SCHEDULE_EMAIL),
		campaignId,
		scheduledTime,
		templateValues,
	}
}

type SendEmailListMessage struct {
	MessagingBase
	ListName string
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
