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
