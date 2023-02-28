package shared_types

const (
	WORK_TYPE_CREATE_GIF = "CREATE_GIF"
)

type CreateGifMessage struct {
	MessagingBase
	EventID           uint
	UserName          string
	UserEmail         string
	UserWalletAddress string
	EmailTemplateId   uint
}

func NewCreateGifMessage(taskId, eventId, emailTemplateId uint, userName, userWalletAddress, userEmail string) *CreateGifMessage {
	return &CreateGifMessage{
		MessagingBase:     *NewMessagingBase(taskId, WORK_TYPE_CREATE_GIF),
		EventID:           eventId,
		EmailTemplateId:   emailTemplateId,
		UserName:          userName,
		UserEmail:         userEmail,
		UserWalletAddress: userWalletAddress,
	}
}
