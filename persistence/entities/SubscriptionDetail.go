package entities

import "time"

type SubscriptionDetail struct {
	Id               uint             `json:"subscription_detail_id" gorm:"primaryKey"`
	SubscriptionId   uint             `json:"subscription_id"`
	Service          string           `json:"subscription_service"`
	Cost             float64          `json:"subscription_cost"`
	Deadline         time.Time        `json:"subscription_deadline" gorm:"type:datetime;default:current_timestamp"`
	StartDate        time.Time        `json:"subscription_start_date" gorm:"type:datetime;default:current_timestamp"`
	PaymentFrequency string           `json:"subscription_payment_frequency"`
	CategoryId       uint             `json:"category_id"`
	Category         Category         `gorm:"foreignKey:CategoryId;references:Id"`
	UserSubscription UserSubscription `gorm:"foreignKey:SubscriptionId;references:Id"`
}

func (s *SubscriptionDetail) DBTableName() string {
	return "subscription_details"
}

func (s *SubscriptionDetail) EntityName() string {
	return "SubscriptionDetail"
}

func (s *SubscriptionDetail) EntityFields() []string {
	return []string{"Id", "SubscriptionId", "Service", "Cost", "Deadline", "StartDate", "PaymentFrequency"}
}
