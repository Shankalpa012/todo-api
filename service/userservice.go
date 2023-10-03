package service

import (
	"fmt"
	"os"
	"todo/bootstraps"
	"todo/model"

	"github.com/golang-jwt/jwt"
)

type UserService struct {
	*bootstraps.Database
}

func NewUserService(DB *bootstraps.Database) *UserService {
	return &UserService{DB}
}

func (db UserService) UserCreate(user *model.User) error {
	return db.Create(&user).Error
}

func (db UserService) ValidateUser(data *model.UserLogin) (user model.User, err error) {
	return user, db.Where("email = ?", data.Email).Find(&user).Error
}

func (db UserService) ValidateToken(tokenString string) (*jwt.Token, error) {
	// Parse and validate the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method")
		}
		return []byte(os.Getenv("SECRETE")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return token, nil
}
