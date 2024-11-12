package DTO

import (
	"proyecto_go/persistence/entities"
	"time"
)

type SubscriptionDetailDTO struct {
	Service          string    `json:"service"`
	Cost             float64   `json:"cost"`
	PaymentFrequency string    `json:"paymentFrequency"`
	Deadline         time.Time `json:"deadline"`
	StartDate        time.Time `json:"startTime"`
	CategoryId       uint      `json:"categoryId"`
}

func (s *SubscriptionDetailDTO) ToEntity() entities.SubscriptionDetail {
	return entities.SubscriptionDetail{
		Service:          s.Service,
		Cost:             s.Cost,
		PaymentFrequency: s.PaymentFrequency,
		Deadline:         s.Deadline,
		StartDate:        s.StartDate,
		CategoryId:       s.CategoryId,
	}
}
