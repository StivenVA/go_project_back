package services

import (
	"errors"
	"proyecto_go/DTO/response"
	"proyecto_go/persistence/entities"
	"proyecto_go/persistence/repositories"
)

func UpdatePaymentStatus(idToken string, response response.PaymentResponse) (any, error) {

	sub, err := ExtractSubClaim(idToken)

	var payment = repositories.FindPaymentById(response.Id)

	if err != nil {
		return nil, err
	}

	var user = repositories.FindUserBySub(sub)

	var payments = repositories.GetPaymentsBySub(user.UserSub)

	if !containsPayment(payments, payment) {
		return nil, errors.New("Payment not found")
	}

	payment.PaymentStatus = entities.StatusPayment(response.Status)

	repositories.UpdatePaymentStatus(payment)

	return "Payment updated", nil
}

func containsPayment(payments []entities.Payment, payment entities.Payment) bool {
	for _, p := range payments {
		if p.Id == payment.Id && p.SubscriptionId == payment.SubscriptionId && p.PaymentStatus == payment.PaymentStatus {
			return true
		}
	}
	return false
}
