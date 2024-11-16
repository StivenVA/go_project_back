package repositories

import (
	"proyecto_go/persistence"
	"proyecto_go/persistence/entities"
)

func FindUserBySub(sub string) entities.User {

	db := persistence.GetConnection()
	var user entities.User
	db.Where("user_sub = ?", sub).First(&user)
	return user

}

func FindUserByEmail(email string) entities.User {

	db := persistence.GetConnection()
	var user entities.User
	db.Where("email = ?", email).First(&user)
	return user
}

func CreateUser(user entities.User) (entities.User, error) {

	db := persistence.GetConnection()
	err := db.Create(&user)

	if err.Error != nil {
		return entities.User{}, err.Error
	}

	erro := CreateSubscription(user.Id)
	if erro != nil {
		return entities.User{}, erro
	}

	return user, nil
	
}
