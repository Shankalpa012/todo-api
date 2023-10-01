package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" unique:"true"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email" unique:"true"`
	Password string `json:"password"`
}

type ProfileData struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
