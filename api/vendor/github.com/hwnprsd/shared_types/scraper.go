package shared_types

import "time"

const WORK_TYPE_SCRAPE_URL = "SCRAPE_URL"
const WORK_TYPE_SUMMARIZE_BLOG = "SUMMARIZE_BLOG"
const WORK_TYPE_UPDATE_RSS = "UPDATE_RSS"
const WORK_TYPE_CREATE_RSS_NEWSLETTER = "CREATE_RSS_NEWSLETTER"
const WORK_TYPE_SEND_RSS_NEWSLETTER = "SEND_RSS_NEWSLETTER"

type ScrapeUrlMessage struct {
	MessagingBase
	Url   string
	Flags string
}

func NewCreateScrapeUrlMessage(taskId uint, url, flags string) *ScrapeUrlMessage {
	return &ScrapeUrlMessage{
		*NewMessagingBase(taskId, WORK_TYPE_SCRAPE_URL),
		url,
		flags,
	}
}

type SummarizeBlogMessage struct {
	MessagingBase
	ScrapeData string
}

func NewSummarizeBlogMessage(taskId uint, data string) *SummarizeBlogMessage {
	return &SummarizeBlogMessage{
		*NewMessagingBase(taskId, WORK_TYPE_SUMMARIZE_BLOG),
		data,
	}
}

type UpdateRssData struct {
	MessagingBase
}

func NewUpdateRssData(taskId uint) *UpdateRssData {
	return &UpdateRssData{
		*NewMessagingBase(taskId, WORK_TYPE_UPDATE_RSS),
	}
}

type CreateRssNewsletterMessage struct {
	MessagingBase
	Tag  string
	Date time.Time
}

func NewCreateRssNewsletterMessage(taskId uint, tag string, newsletterPublishedDate time.Time) *CreateRssNewsletterMessage {
	return &CreateRssNewsletterMessage{
		*NewMessagingBase(taskId, WORK_TYPE_CREATE_RSS_NEWSLETTER),
		tag, newsletterPublishedDate,
	}
}

type SendRssNewsletterMessage struct {
	MessagingBase
	Tag           string
	Date          time.Time
	CampaignId    uint
	ScheduledTime time.Time
}

func NewSendRssNewsletterMessage(taskId, campaignId uint, tag string, newsletterPublishedDate, scheduledTime time.Time) *SendRssNewsletterMessage {
	return &SendRssNewsletterMessage{
		*NewMessagingBase(taskId, WORK_TYPE_SEND_RSS_NEWSLETTER),
		tag, newsletterPublishedDate, campaignId, scheduledTime,
	}
}
