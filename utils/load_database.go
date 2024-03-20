package utils

import (
	"blog/db"
	"blog/model"
)

func LoadDatabase() {
	db.Connect()

	db.Database.AutoMigrate(&model.User{})
	db.Database.AutoMigrate(&model.Blog{})
	db.Database.AutoMigrate(&model.Comment{})
}