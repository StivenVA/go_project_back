package entities

type StatusPayment string

const (
	Pending   StatusPayment = "PENDING"
	Success   StatusPayment = "SUCCESS"
	Failed    StatusPayment = "FAILED"
	Cancelled StatusPayment = "CANCELLED"
	Error     StatusPayment = "ERROR"
	OVERDUE   StatusPayment = "OVERDUE"
)
