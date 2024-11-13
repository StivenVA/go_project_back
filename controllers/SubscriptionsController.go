package controllers

import (
	"net/http"
	"proyecto_go/DTO"
	"proyecto_go/DTO/response"
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

func GetSubscriptions() (string, http.HandlerFunc) {
	return "/getSubscriptions", func(w http.ResponseWriter, r *http.Request) {

		var idToken = r.Header.Get("IdToken")

		if r.Method == http.MethodGet {
			resp, erro := services.GetSubscriptions(idToken)
			responseManager(w, resp, erro)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
}

func CreateCategory() (string, http.HandlerFunc) {
	return "/categories", func(w http.ResponseWriter, r *http.Request) {

		var category response.CategoryDTO

		idToken := r.Header.Get("IdToken")

		decodeRequest(w, r, &category)

		if r.Method == http.MethodPost {
			resp, erro := services.CreateCategory(idToken, category)
			responseManager(w, resp, erro)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
}

func GetCategories() (string, http.HandlerFunc) {
	return "/getCategories", func(w http.ResponseWriter, r *http.Request) {

		var idToken = r.Header.Get("IdToken")

		if r.Method == http.MethodGet {
			resp, erro := services.GetCategories(idToken)
			responseManager(w, resp, erro)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
}
