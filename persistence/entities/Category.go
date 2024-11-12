package entities

type Category struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserId int    `json:"user_id"`
	User   User   `gorm:"foreignKey:UserId;references:Id"`
}

func (c Category) DBTableName() string {
	return "categories"
}

func (c Category) EntityName() string {
	return "Category"
}

func (c Category) EntityFields() []string {
	return []string{"ID", "Name", "UserId"}
}
