package entities

import "time"

type Payment struct {
	Id                 uint                 `json:"payment_id" gorm:"primaryKey"`
	SubscriptionDetail []SubscriptionDetail `json:"subscription_detail" gorm:"foreignKey:SubscriptionId"`
	Date               time.Time            `json:"payment_date" gorm:"type:datetime;default:current_timestamp"`
	Amount             float64              `json:"payment_amount"`
	Status             StatusPayment        `json:"payment_status"`
}

func (p *Payment) DBTableName() string {
	return "payments"
}

func (p *Payment) EntityName() string {
	return "Payment"
}

func (p *Payment) EntityFields() []string {
	return []string{"Id","SubscriptionDetail","Date","Amount","Status"}
}
