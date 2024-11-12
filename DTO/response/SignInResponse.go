package response

type SignInResponse struct {
	User  any    `json:"user"`
	Token string `json:"token"`
}
