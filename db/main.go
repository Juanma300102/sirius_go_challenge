package db

import (
	"challenge/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func GetConnection() *gorm.DB {
	if db == nil {
		db, err = gorm.Open(sqlite.Open("test.sqlite"))
		if err != nil {
			panic("Failed to connect to the database")
		}
	}
	return db
}

func Migrate() {
	db = GetConnection()
	db.AutoMigrate(&models.User{})
}