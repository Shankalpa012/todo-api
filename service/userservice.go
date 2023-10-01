package service

import (
	"todo/bootstraps"
	"todo/model"
)

func UserCreate(user *model.User) error {
	return bootstraps.DB.Create(&user).Error
}

func ValidateUser(data *model.UserLogin) (user model.User, err error) {
	return user, bootstraps.DB.Where("email = ?", data.Email).Find(&user).Error
}
