package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"todo/model"
	"todo/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(ctx *gin.Context) {
	//define the data
	var data model.User

	//validate the data
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong Data Was Provided",
		})
		return
	}

	//hash the password
	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)

	user := model.User{
		Email:    data.Email,
		Password: string(password),
	}

	//save to the database
	err := service.UserCreate(&user)
	if err != nil {
		fmt.Println("Error while saving data to database")
	}

	//return response
	ctx.JSON(200, gin.H{
		"message": "User Created Successfully",
	})
}

func Login(ctx *gin.Context) {
	//get the data from the body

	var data model.UserLogin

	//validate the request data
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid data was provided",
		})
		return
	}

	// validate if the user exists or not
	user, err := service.ValidateUser(&data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "System Query Error",
		})
		return
	}

	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "No Record Found",
		})
		return
	}

	// check if the email and password are correct or not
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Incorrect Password or User Name",
		})
		return
	}

	//generate the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	//providing secrete key to the token
	tokenstring, err := token.SignedString([]byte(os.Getenv("SECRETE")))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Generate the token",
			"err":     err.Error(),
		})
		return
	}

	//return the success with token
	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenstring,
		"email": user.Email,
		"id":    user.ID,
	})
}
