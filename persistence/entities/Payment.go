package entities

import (
	"proyecto_go/DTO/response"
	"proyecto_go/persistence/repositories"
	"time"
)

type Payment struct {
	Id                 uint               `json:"payment_id" gorm:"primaryKey"`
	PaymentDate        time.Time          `json:"payment_date" gorm:"type:datetime;default:current_timestamp"`
	Amount             float64            `json:"payment_amount"`
	PaymentStatus      StatusPayment      `json:"payment_status" gorm:"type:enum('PENDING', 'SUCCESS', 'FAILED', 'CANCELLED', 'ERROR', 'OVERDUE');default:'PENDING'"`
	SubscriptionId     uint               `json:"subscription_id"`
	SubscriptionDetail SubscriptionDetail `gorm:"foreignKey:SubscriptionId;references:Id"`
}

func (p *Payment) DBTableName() string {
	return "payments"
}

func (p *Payment) EntityName() string {
	return "Payment"
}

func (p *Payment) EntityFields() []string {
	return []string{"Id", "SubscriptionDetail", "Date", "Amount", "Status"}
}

func (p *Payment) PaymentToDTO() response.PaymentResponse {

	var subscription = repositories.FindSubscriptionDetailById(p.SubscriptionId)

	return response.PaymentResponse{
		Id:      p.Id,
		Date:    p.PaymentDate.Format("2006-01-02"),
		Amount:  p.Amount,
		Status:  string(p.PaymentStatus),
		Service: subscription.Service,
	}

}

func PaymentToDTOList(payments []Payment) []response.PaymentResponse {
	var paymentsDTO []response.PaymentResponse
	for _, payment := range payments {
		paymentsDTO = append(paymentsDTO, payment.PaymentToDTO())
	}
	return paymentsDTO
}
