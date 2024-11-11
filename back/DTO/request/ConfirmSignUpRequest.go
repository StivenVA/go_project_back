package request

type ConfirmSignUpRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}
