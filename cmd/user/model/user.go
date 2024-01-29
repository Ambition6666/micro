package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"user_name"`
	Password string `json:"-"`
	Email    string `json:"email"`
}

func NewUser(name string, pwd string, email string) *User {
	return &User{
		UserName: name,
		Password: pwd,
		Email:    email,
	}
}
