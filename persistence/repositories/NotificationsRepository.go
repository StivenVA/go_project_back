package repositories

import (
	"proyecto_go/persistence"
	"proyecto_go/persistence/entities"
	"time"
)

func CreateNotification(notification entities.Notifications) (entities.Notifications, error) {

	db := persistence.GetConnection()
	err := db.Create(&notification)

	if err.Error != nil {
		return notification, err.Error
	}

	return notification, nil
}

func GetNotificationsByUserSub(sub string) []entities.Notifications {

	db := persistence.GetConnection()

	var user entities.User

	db.Where("user_sub = ?", sub).First(&user)

	var notifications []entities.Notifications

	db.Where("user_id = ? and notification_date = "+time.Now().String(), user.Id).Find(&notifications)

	return notifications
}
