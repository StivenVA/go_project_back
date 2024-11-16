package services

import (
	"proyecto_go/persistence/entities"
	"proyecto_go/persistence/repositories"
)

func GetNotifications(idToken string) (any, error) {

	sub, err := ExtractSubClaim(idToken)

	if err != nil {
		return nil, err
	}

	var user = repositories.FindUserBySub(sub)

	var notifications = repositories.GetNotificationsByUserSub(user.UserSub)

	return entities.NotificationToDTOList(notifications), nil

}
