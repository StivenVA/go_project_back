package persistence

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	dsn := "admin:goland123@tcp(goland.c3as48eygymf.us-east-1.rds.amazonaws.com:3306)/go_project?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}
