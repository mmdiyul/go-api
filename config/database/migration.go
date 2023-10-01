package database

import (
	"go-api/app/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.User{})
}
