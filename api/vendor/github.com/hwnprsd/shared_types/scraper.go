package shared_types

const WORK_TYPE_SCRAPE_URL = "SCRAPE_URL"
const WORK_TYPE_SUMMARIZE_BLOG = "SCRAPE_URL"

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
