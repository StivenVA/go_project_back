package request

type AuthUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
