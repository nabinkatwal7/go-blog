package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Body   string `gorm:"type:text; not null;" json:"body"`
	Blog   Blog
	BlogID uint
	User   User
	UserID uint
}