package repositories

import (
	"proyecto_go/persistence"
	"proyecto_go/persistence/entities"
)

func CreateCategory(category entities.Category) (entities.Category, error) {

	db := persistence.GetConnection()
	err := db.Create(&category)

	if err.Error != nil {
		return category, err.Error
	}

	return category, nil
}

func GetCategoriesByUserSub(sub string) []entities.Category {

	db := persistence.GetConnection()

	var user entities.User

	db.Where("user_sub = ?", sub).First(&user)

	var categories []entities.Category

	db.Where("user_id = ?", user.Id).Find(&categories)

	return categories
}

func FindCategoryById(id uint) entities.Category {

	db := persistence.GetConnection()

	var category entities.Category

	db.Where("id = ?", id).First(&category)

	return category
}

func DeleteCategoryById(id uint) error {

	db := persistence.GetConnection()

	var category entities.Category

	db.Where("id = ?", id).First(&category)

	err := db.Delete(&category)

	if err.Error != nil {
		return err.Error
	}

	return nil
}
