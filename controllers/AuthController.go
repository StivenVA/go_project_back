package controllers

import (
	"net/http"
	"proyecto_go/DTO/request"
	"proyecto_go/services"
)

func LoginHandler() (string, http.HandlerFunc) {
	return "/login", func(w http.ResponseWriter, r *http.Request) {

		var auth request.AuthUser

		decodeRequest(w, r, &auth)

		if r.Method == http.MethodPost {
			resp, erro := services.Login(auth)
			responseManager(w, resp, erro)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
}

func RegisterHandler() (string, http.HandlerFunc) {
	return "/signup", func(w http.ResponseWriter, r *http.Request) {

		var regist request.RegisterRequest

		decodeRequest(w, r, &regist)

		if r.Method == http.MethodPost {
			resp, erro := services.SignUp(regist)
			responseManager(w, resp, erro)

		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
}

func ConfirmEmail() (string, http.HandlerFunc) {
	return "/confirmEmail", func(w http.ResponseWriter, r *http.Request) {

		var confirm request.ConfirmSignUpRequest

		decodeRequest(w, r, &confirm)

		if r.Method == http.MethodPost {
			resp, erro := services.ConfirmSignUp(confirm)
			responseManager(w, resp, erro)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
}

func SocialLogin() (string, http.HandlerFunc) {
	return "/socialLogin", func(w http.ResponseWriter, r *http.Request) {

		var auth request.SocialLogin

		decodeRequest(w, r, &auth)

		if r.Method == http.MethodPost {
			resp, erro := services.SocialLogin(auth)
			responseManager(w, resp, erro)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
}

func ResendConfirmationCode() (string, http.HandlerFunc) {
	return "/resendConfirmationCode", func(w http.ResponseWriter, r *http.Request) {

		var resend request.ResendConfirmationCodeRequest

		decodeRequest(w, r, &resend)

		if r.Method == http.MethodPost {
			resp, erro := services.ResendConfirmationCode(resend)
			responseManager(w, resp, erro)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
}
