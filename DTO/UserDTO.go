package DTO

import "proyecto_go/persistence/entities"

type UserDTO struct {
	Sub      string `json:"sub"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Id       string `json:"id"`
	Phone    string `json:"phone"`
}

func UserDTOToEntity(userDTO UserDTO) entities.User {
	return entities.User{
		UserSub: userDTO.Sub,
		Email:   userDTO.Email,
		Name:    userDTO.Name,
		Phone:   userDTO.Phone,
	}
}
