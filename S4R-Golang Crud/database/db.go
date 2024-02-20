package database

import (
	"crud-demo/app/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


var db *gorm.DB

func InitDB() {
    database, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }
    database.AutoMigrate(&models.User{})
    db = database
}

func GetDB() *gorm.DB {
    return db
}