package controllers

import (
	"net/http"
	"proyecto_go/services"
)

func GenerateNotification() (string, http.HandlerFunc) {
	return "/GenerateNotification", func(w http.ResponseWriter, r *http.Request) {

		var idToken = r.Header.Get("IdToken")

		if r.Method == http.MethodGet {
			resp, erro := services.GetSubscriptions(idToken)
			responseManager(w, resp, erro)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}

	}
}
