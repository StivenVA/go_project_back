package controllers

import (
	"net/http"
	"proyecto_go/services"
)

func LoginHandler() (string, http.HandlerFunc) {
	return "/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			services.Login(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
}

func RegisterHandler() (string, http.HandlerFunc) {
	return "/signup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			services.SignUp(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
}

func ConfirmEmail() (string, http.HandlerFunc) {
	return "/confirmEmail", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			services.ConfirmSignUp(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
}

func GetAuthEndPoints() []func() (string, http.HandlerFunc) {

	authEndPoints := []func() (string, http.HandlerFunc){
		LoginHandler,
		RegisterHandler,
		ConfirmEmail,
	}

	return authEndPoints
}
