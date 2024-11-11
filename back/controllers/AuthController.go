package controllers

import (
	"encoding/json"
	"net/http"
	"proyecto_go/DTO/request"
	"proyecto_go/DTO/response"
	"proyecto_go/services"
)

func LoginHandler() (string, http.HandlerFunc) {
	return "/login", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		var auth request.AuthUser

		err := json.NewDecoder(r.Body).Decode(&auth)
		if err != nil {
			json.NewEncoder(w).Encode("Invalid request")
		}

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

		w.Header().Set("Content-Type", "application/json")
		var regist request.RegisterRequest

		err := json.NewDecoder(r.Body).Decode(&regist)
		if err != nil {
			json.NewEncoder(w).Encode("Invalid request")
			return
		}

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

		w.Header().Set("Content-Type", "application/json")
		var confirm request.ConfirmSignUpRequest

		err := json.NewDecoder(r.Body).Decode(&confirm)
		if err != nil {
			json.NewEncoder(w).Encode("Invalid request")
			return
		}

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

		w.Header().Set("Content-Type", "application/json")
		var auth request.SocialLogin

		err := json.NewDecoder(r.Body).Decode(&auth)
		if err != nil {
			json.NewEncoder(w).Encode("Invalid request")
		}

		if r.Method == http.MethodPost {
			resp, erro := services.SocialLogin(auth)
			responseManager(w, resp, erro)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
}

func GetAuthEndPoints() []func() (string, http.HandlerFunc) {

	return []func() (string, http.HandlerFunc){
		LoginHandler,
		RegisterHandler,
		ConfirmEmail,
	}
}

func responseManager(w http.ResponseWriter, resp any, erro error) {

	if erro != nil {
		json.NewEncoder(w).Encode(response.ResponseEntity{
			Message: erro.Error(),
			Status:  http.StatusInternalServerError,
		})
	} else {
		json.NewEncoder(w).Encode(response.ResponseEntity{
			Status: http.StatusOK,
			Data:   resp,
		})
	}

}
