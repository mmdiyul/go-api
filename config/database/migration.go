package database

import (
	"go-api/app/models/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&user.Role{})
	db.AutoMigrate(&user.User{})
}
