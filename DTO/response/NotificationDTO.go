package response

type NotificationDTO struct {
	NotificationId uint   `json:"id"`
	UserId         uint   `json:"user_id"`
	Message        string `json:"message"`
	Status         string `json:"status"`
}
