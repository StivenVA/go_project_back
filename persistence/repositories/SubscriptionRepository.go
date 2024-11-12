package repositories

import (
	"proyecto_go/persistence"
	"proyecto_go/persistence/entities"
)

func CreateSubscription(userId uint) error {

	db := persistence.GetConnection()
	err := db.Create(&entities.UserSubscription{UserId: userId})

	if err.Error != nil {
		return err.Error
	}

	return nil
}

func SaveDetails(subscription entities.SubscriptionDetail) error {

	db := persistence.GetConnection()
	err := db.Create(&subscription)

	if err.Error != nil {
		return err.Error
	}

	return nil
}

func FindSubscriptionByUserSub(sub string) entities.UserSubscription {

	db := persistence.GetConnection()

	var user entities.User

	db.Where("user_sub = ?", sub).First(&user)

	var subscription entities.UserSubscription

	db.Where("user_id = ?", user.Id).First(&subscription)

	db.Where("subscription_id = ?", subscription.Id).First(&subscription.SubscriptionDetail)

	return subscription
}
