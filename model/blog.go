package model

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title string `gorm:"size:255; not null;" json:"title"`
	Body  string `gorm:"type:text; not null;" json:"body"`
	User  User  
	UserID uint
	Comments []Comment
}