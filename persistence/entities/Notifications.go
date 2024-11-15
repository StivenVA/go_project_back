package entities

import "time"

type Notifications struct {
	Id                 uint                 `json:"notification_id" gorm:"primaryKey"`
	UserId             uint                 `json:"user_id"` // Definimos la relación abajo, no en este campo
	SubscriptionDetail []SubscriptionDetail `json:"subscription_detail" gorm:"foreignKey:SubscriptionId;references:Id"`
	Date               time.Time            `json:"notification_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	Status             string               `json:"notification_status" gorm:"not null;enum:('PENDING','SENT','READ')"`
	Message            string               `json:"notification_message" gorm:"not null"`
	User               User                 `gorm:"foreignKey:UserId;references:Id"` // Relación con la tabla Users
}

func (n *Notifications) DBTableName() string {
	return "notifications"
}

func (n *Notifications) EntityName() string {
	return "Notifications"
}

func (n *Notifications) EntityFields() []string {
	return []string{"Id", "UserId", "SubscriptionDetail", "Date", "Status", "Message"}
}
