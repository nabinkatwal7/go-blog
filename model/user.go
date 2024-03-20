package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:255; not null;" json:"username"`
	Email    string `gorm:"size:255; not null; unique" json:"email"`
	Password string `gorm:"size:255; not null;" json:"-"`
	ConfirmPassword string `gorm:"size:255; not null;" json:"-"`
	Blogs    []Blog
	Comments []Comment

}