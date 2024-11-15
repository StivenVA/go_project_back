package response

type PaymentResponse struct {
	Id      uint    `json:"id"`
	Date    string  `json:"date"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
	Service string  `json:"service"`
}
