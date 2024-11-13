package services

import (
	"errors"
	"proyecto_go/DTO"
	"proyecto_go/DTO/response"
	"proyecto_go/persistence/entities"
	"proyecto_go/persistence/repositories"
)

func CreateSubscription(idToken string, detail DTO.SubscriptionDetailDTO) (any, error) {

	sub, err := ExtractSubClaim(idToken)

	if err != nil {
		return "Error extracting sub claim", err
	}

	user := repositories.FindUserBySub(sub)

	if user.Id == 0 {
		return "User not found", errors.New("User not found")
	}

	subscription := repositories.FindSubscriptionByUserSub(sub)

	detailEntity := detail.ToEntity()

	subscription.User = user
	detailEntity.SubscriptionId = subscription.Id
	detailEntity.UserSubscription = subscription

	err = repositories.SaveDetails(detailEntity)

	if err != nil {
		return "Error saving subscription details", err
	}

	return detail, nil
}

func GetSubscriptions(idToken string) (any, error) {

	sub, err := ExtractSubClaim(idToken)

	if err != nil {
		return "Error extracting sub claim", err
	}

	user := repositories.FindUserBySub(sub)

	if user.Id == 0 {
		return "User not found", errors.New("User not found")
	}

	subscription := repositories.FindSubscriptionByUserSub(sub)

	subscriptionResponse := response.SubscriptionResponse{
		Id:            subscription.Id,
		UserId:        user.Id,
		Subscriptions: entities.ToDTOList(subscription.SubscriptionDetail),
		UserName:      user.Name,
	}

	if subscription.Id == 0 {
		return "Subscription not found", errors.New("Subscription not found")
	}

	return subscriptionResponse, nil
}
