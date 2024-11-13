package controllers

import (
	"encoding/json"
	"net/http"
	"proyecto_go/DTO/response"
)

func responseManager(w http.ResponseWriter, resp any, erro error) {

	w.Header().Set("Content-Type", "application/json")

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(response.ResponseEntity{
			Message: erro.Error(),
			Status:  http.StatusBadRequest,
		})
	} else {
		json.NewEncoder(w).Encode(response.ResponseEntity{
			Status: http.StatusOK,
			Data:   resp,
		})
	}

}

func decodeRequest(w http.ResponseWriter, r *http.Request, data any) {

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid request")
		return
	}

}

func GetEndPoints() []func() (string, http.HandlerFunc) {

	return []func() (string, http.HandlerFunc){
		LoginHandler,
		RegisterHandler,
		ConfirmEmail,
		SocialLogin,
		ResendConfirmationCode,
		CreateSubscription,
		GetSubscriptions,
	}
}
