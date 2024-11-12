package entities

type UserSubscription struct {
	Id                 uint                 `json:"subscription_id" gorm:"primaryKey"`
	UserId             uint                 `json:"user_id"`
	SubscriptionDetail []SubscriptionDetail `json:"subscription_detail" gorm:"foreignKey:SubscriptionId"`
	User               User                 `gorm:"foreignKey:UserId;references:Id"`
}

func (us *UserSubscription) DBTableName() string {
	return "user_subscriptions"
}

func (us *UserSubscription) EntityName() string {
	return "UserSubscription"
}

func (us *UserSubscription) EntityFields() []string {
	return []string{"Id", "UserId", "SubscriptionDetail"}
}
