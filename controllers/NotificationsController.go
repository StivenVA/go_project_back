package controllers

import (
	"net/http"
	"proyecto_go/services"
)

func GetUserNotifications() (string, http.HandlerFunc) {
	return "/notifications", func(w http.ResponseWriter, r *http.Request) {

		var idToken = r.Header.Get("IdToken")

		if r.Method == http.MethodGet {
			resp, erro := services.GetNotifications(idToken)
			responseManager(w, resp, erro)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
}
