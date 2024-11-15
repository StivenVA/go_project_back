package repositories

import (
	"proyecto_go/persistence"
	"proyecto_go/persistence/entities"
)

func CreatePayment(payment entities.Payment) (entities.Payment, error) {

	db := persistence.GetConnection()
	err := db.Create(&payment)

	if err.Error != nil {
		return payment, err.Error
	}

	return payment, nil
}

func UpdatePaymentStatus(payment entities.Payment) {
	db := persistence.GetConnection()

	db.Model(payment).Update("payment_status", payment.PaymentStatus)

	if payment.PaymentStatus == entities.Success {

		var subscription entities.SubscriptionDetail

		switch payment.SubscriptionDetail.PaymentFrequency {
		case entities.Daily:
			db.Raw("SELECT * FROM subscription_details WHERE id = ?", payment.SubscriptionId).Scan(&subscription)
			db.Model(subscription).Update("deadline", subscription.Deadline.AddDate(0, 0, 1))
		case entities.Weekly:
			db.Raw("SELECT * FROM subscription_details WHERE id = ?", payment.SubscriptionId).Scan(&subscription)
			db.Model(subscription).Update("deadline", subscription.Deadline.AddDate(0, 0, 7))
		case entities.Monthly:
			db.Raw("SELECT * FROM subscription_details WHERE id = ?", payment.SubscriptionId).Scan(&subscription)
			db.Model(subscription).Update("deadline", subscription.Deadline.AddDate(0, 1, 0))
		case entities.Yearly:
			db.Raw("SELECT * FROM subscription_details WHERE id = ?", payment.SubscriptionId).Scan(&subscription)
			db.Model(subscription).Update("deadline", subscription.Deadline.AddDate(1, 0, 0))
		}

		newPayment := entities.Payment{
			SubscriptionId: subscription.SubscriptionId,
			PaymentStatus:  entities.Pending,
			PaymentDate:    subscription.Deadline,
			Amount:         subscription.Cost,
		}

		CreatePayment(newPayment)
	}

}

func GetPaymentsBySub(sub string) []entities.Payment {
	db := persistence.GetConnection()
	var payments []entities.Payment

	query := "SELECT p.* FROM payments p,subscription_details sd, users u, user_subscriptions us WHERE p.subscription_id = sd.id and sd.subscription_id = us.id and u.id = us.user_id and u.user_sub = ?"

	db.Raw(query, sub).Scan(&payments)

	return payments
}

func FindPaymentById(id uint) entities.Payment {
	db := persistence.GetConnection()
	var payment entities.Payment
	db.Where("id = ?", id).First(&payment)
	return payment
}
