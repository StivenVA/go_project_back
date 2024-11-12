package controllers

import (
	"net/http"
	"proyecto_go/DTO"
	"proyecto_go/services"
)

func CreateSubscription() (string, http.HandlerFunc) {
	return "/subscriptions", func(w http.ResponseWriter, r *http.Request) {

		var subscription DTO.SubscriptionDetailDTO

		var idToken = r.Header.Get("IdToken")

		decodeRequest(w, r, &subscription)

		if r.Method == http.MethodPost {
			resp, erro := services.CreateSubscription(idToken, subscription)
			responseManager(w, resp, erro)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
}
