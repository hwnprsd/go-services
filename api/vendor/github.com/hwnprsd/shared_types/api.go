package shared_types

const (
	WORK_TYPE_API_CALLBACK = "CALLBACK_API"
)

type ApiCallback struct {
	MessagingBase
	TaskID uint
	Status string
	Data   any
}

func NewApiCallbackMessage(taskId uint, status string, data any) *ApiCallback {
	return &ApiCallback{
		MessagingBase: *NewMessagingBase(taskId, WORK_TYPE_API_CALLBACK),
		TaskID:        taskId,
		Status:        status,
		Data:          data,
	}
}
