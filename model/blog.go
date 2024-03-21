package model

import (
	"blog/db"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title string `gorm:"size:255; not null;" json:"title"`
	Body  string `gorm:"type:text; not null;" json:"body"`
	User  User  
	UserID uint
	Comments []Comment
}

func (b *Blog) Save() (*Blog, error) {
	err := db.Database.Create(&b).Error
	if err != nil {
		return &Blog{}, err
	}
	return b, nil
}

func (b *Blog) Update() (*Blog, error) {
	err := db.Database.Save(&b).Error
	if err != nil {
		return &Blog{}, err
	}
	return b, nil
}

func (b *Blog) Delete() error{
	return db.Database.Delete(&b).Error
}