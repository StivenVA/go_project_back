package persistence

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var connection *gorm.DB

func ConnectDB() (*gorm.DB, error) {

	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}

func GetConnection() *gorm.DB {
	return connection
}

func SetConnection(db *gorm.DB) {
	connection = db
}
