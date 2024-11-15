package entities

type PaymentFrequency string

const (
	Daily   PaymentFrequency = "DAILY"
	Weekly  PaymentFrequency = "WEEKLY"
	Monthly PaymentFrequency = "MONTHLY"
	Yearly  PaymentFrequency = "YEARLY"
)
