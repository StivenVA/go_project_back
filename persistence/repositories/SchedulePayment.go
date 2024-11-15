package repositories

import (
	"proyecto_go/persistence"
	"proyecto_go/persistence/entities"
)

func FindNextPaymentsByUserSub(sub string) []entities.Payment {

	db := persistence.GetConnection()
	var payments []entities.Payment

	query := "select p.* from payments p,subscriptions s,users u,user_subscriptions us where p.subscription_id = s.id and s.subscription_id = us.id and u.id = us.user_id and u.user_sub = ? and p.status = 'PENDING'"

	db.Raw(query, sub).Scan(&payments)

	return payments
}
