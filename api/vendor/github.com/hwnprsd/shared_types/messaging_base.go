package shared_types

type MessagingBase struct {
	TaskId   uint
	WorkType string
}

func NewMessagingBase(TaskId uint, WorkType string) *MessagingBase {
	return &MessagingBase{TaskId, WorkType}
}
