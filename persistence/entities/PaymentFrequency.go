package entities

type PaymentFrequency string

const (
	Weekly  PaymentFrequency = "WEEKLY"
	Monthly PaymentFrequency = "MONTHLY"
	Yearly  PaymentFrequency = "YEARLY"
)
