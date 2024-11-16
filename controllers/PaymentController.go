package controllers

import (
	"net/http"
	"proyecto_go/DTO/response"
	"proyecto_go/services"
)

func UpdatePayment() (string, http.HandlerFunc) {
	return "UpdatePayment", func(w http.ResponseWriter, r *http.Request) {

		idToken := r.Header.Get("IdToken")

		var payment response.PaymentResponse

		decodeRequest(w, r, &payment)

		if r.Method == http.MethodPut {
			resp, erro := services.UpdatePaymentStatus(idToken, payment)
			responseManager(w, resp, erro)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}

	}
}

func GetPayments() (string, http.HandlerFunc) {
	return "GetPayments", func(w http.ResponseWriter, r *http.Request) {

		idToken := r.Header.Get("IdToken")

		if r.Method == http.MethodGet {
			resp, erro := services.GetPayments(idToken)
			responseManager(w, resp, erro)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}

	}
}

func GetNextPayments() (string, http.HandlerFunc) {
	return "GetNextPayments", func(w http.ResponseWriter, r *http.Request) {

		idToken := r.Header.Get("IdToken")

		if r.Method == http.MethodGet {
			resp, erro := services.GetNextPayments(idToken)
			responseManager(w, resp, erro)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}

	}
}
