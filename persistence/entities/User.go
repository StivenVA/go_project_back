package entities

type User struct {
	Id      uint   `json:"user_id" gorm:"primaryKey"`
	Name    string `json:"user_name" gorm:"not null"`
	Email   string `json:"user_email" gorm:"not null"`
	UserSub string `json:"user_sub" gorm:"not null"`
	Phone   string `json:"user_phone" gorm:"not null"`
}

func (u *User) DBTableName() string {
	return "users"
}

func (u *User) EntityName() string {
	return "User"
}

func (u *User) EntityFields() []string {
	return []string{"Id", "Name", "Email", "UserSub", "Phone"}
}
