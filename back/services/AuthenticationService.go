package services

import (
	"proyecto_go/DTO/request"
	"proyecto_go/DTO/response"
	"proyecto_go/persistence/entities"
	"proyecto_go/persistence/repositories"
)

func Login(user request.AuthUser) (response.SignInResponse, error) {

	authResponse, err := CognitoLogin(user)

	if err != nil {
		return response.SignInResponse{}, err
	}

	token := *authResponse.IdToken

	userEntity := repositories.FindUserByEmail(user.Email)

	return response.SignInResponse{User: userEntity, Token: token}, nil

}

func SignUp(user request.RegisterRequest) (string, error) {
	success, err := CognitoSignUp(user.Email, user.Password)

	if err != nil {
		return "", err
	}
	userEntity := entities.User{
		Email:   user.Email,
		Name:    user.Name,
		Phone:   user.Phone,
		UserSub: *success.UserSub,
	}

	_, err = repositories.CreateUser(userEntity)

	if err != nil {
		return "", err
	}

	return "User created successfully, please confirm your email", nil
}

func ConfirmSignUp(confirm request.ConfirmSignUpRequest) (string, error) {
	err := CognitoConfirmSignUp(confirm)

	if err != nil {
		return "", err
	}

	return "User confirmed successfully", nil

}
