package shared_types

const WORK_TYPE_PDF_PARSE_CV = "PDF_PARSE_CV"

type PdfParseCVMessage struct {
	MessagingBase
	Url string
}

func NewPdfParseMessage(taskId uint, url string) *PdfParseCVMessage {
	return &PdfParseCVMessage{
		*NewMessagingBase(taskId, WORK_TYPE_PDF_PARSE_CV),
		url,
	}
}
