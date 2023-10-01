package service

import (
	"fmt"
	"todo/bootstraps"
	"todo/model"

	"github.com/gin-gonic/gin"
)

func Create(todo *model.Todo) error {
	return bootstraps.DB.Create(&todo).Error
}

func Update(id string, todo map[string]interface{}) error {
	return bootstraps.DB.Model(&model.Todo{}).Where("id = ?", id).Updates(todo).Error
}

func GetAll() (map[string]interface{}, error) {
	var todos []model.Todo

	err := bootstraps.DB.Find(&todos).Error
	if err != nil {
		fmt.Println("Error while getting")
		return nil, err
	}

	return gin.H{"data": todos}, nil
}

// get todo for a specific user
func GetUserTodo(userId string) (map[string]interface{}, error) {
	var todos []model.Todo

	err := bootstraps.DB.Where("user_id = ?", userId).Find(&todos).Error
	if err != nil {
		fmt.Println("Error while getting")
		return nil, err
	}

	return gin.H{"data": todos}, nil
}

func GetById(id string) (todo *model.Todo, err error) {
	return todo, bootstraps.DB.First(&todo, "id = ?", id).Error
}
