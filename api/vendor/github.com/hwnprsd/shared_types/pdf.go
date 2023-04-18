package shared_types

const WORK_TYPE_PDF_PARSE_CV = "PDF_PARSE_CV"
const WORK_TYPE_PDF_PARSE_CV_BYTES = "PDF_PARSE_CV_BYTES"

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

type PdfParseCVBytesMessage struct {
	MessagingBase
	B []byte
}

func NewPdfParseBytesMessage(taskId uint, b []byte) *PdfParseCVBytesMessage {
	return &PdfParseCVBytesMessage{
		*NewMessagingBase(taskId, WORK_TYPE_PDF_PARSE_CV_BYTES),
		b,
	}
}
