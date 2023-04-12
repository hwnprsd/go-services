package shared_types

const WORK_TYPE_SCRAPE_URL = "SCRAPE_URL"
const WORK_TYPE_SUMMARIZE_BLOG = "SUMMARIZE_BLOG"
const WORK_TYPE_UPDATE_RSS = "UPDATE_RSS"

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
