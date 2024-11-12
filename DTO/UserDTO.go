package DTO

type UserDTO struct {
	Sub      string `json:"sub"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Id       string `json:"id"`
	Phone    string `json:"phone"`
}
