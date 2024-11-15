package response

import (
	"proyecto_go/persistence/entities"
	"time"
)

type SubscriptionDTO struct {
	Service          string                    `json:"service"`
	Cost             float64                   `json:"cost"`
	PaymentFrequency entities.PaymentFrequency `json:"paymentFrequency"`
	Deadline         time.Time                 `json:"deadline"`
	StartDate        time.Time                 `json:"startTime"`
	Category         CategoryDTO               `json:"category"`
}

type SubscriptionResponse struct {
	Id            uint              `json:"subscription_id"`
	Subscriptions []SubscriptionDTO `json:"subscriptions"`
	UserName      string            `json:"userName"`
	UserId        uint              `json:"userId"`
}
