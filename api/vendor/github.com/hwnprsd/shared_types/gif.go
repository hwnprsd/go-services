package shared_types

const (
	WORK_TYPE_CREATE_GIF = "CREATE_GIF"
)

// 	// Inputs

type CreateGifMessage struct {
	MessagingBase
	EventID           uint
	UserName          string
	UserEmail         string
	UserWalletAddress string
	EmailTemplateId   uint
}

func NewCreateGifMessage(eventId, emailTemplateId uint, userName, userWalletAddress, userEmail string) *CreateGifMessage {
	return &CreateGifMessage{
		MessagingBase: MessagingBase{
			WorkType: WORK_TYPE_CREATE_GIF,
		},
		EventID:           eventId,
		EmailTemplateId:   emailTemplateId,
		UserName:          userName,
		UserEmail:         userEmail,
		UserWalletAddress: userWalletAddress,
	}
}
