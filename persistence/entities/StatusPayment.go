package entities

type StatusPayment int

const (
	Pending StatusPayment = iota
	Success
	Failed
	Cancelled
	Error
)
