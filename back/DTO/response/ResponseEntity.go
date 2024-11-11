package response

type ResponseEntity struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    interface{}
}
