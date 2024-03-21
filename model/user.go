package model

import (
	"blog/db"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255; not null;" json:"username"`
	Email    string `gorm:"size:255; not null; unique" json:"email"`
	Password string `gorm:"size:255; not null;" json:"-"`
	ConfirmPassword string `gorm:"size:255; not null;" json:"-"`
	Blogs    []Blog
	Comments []Comment

}

func (user *User) Save() (*User, error) {
	err := db.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.ConfirmPassword = string(passwordHash)
	return nil
}