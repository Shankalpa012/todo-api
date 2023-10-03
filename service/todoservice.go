package service

import (
	"fmt"
	"todo/bootstraps"
	"todo/model"

	"github.com/gin-gonic/gin"
)

type TodoService struct {
	*bootstraps.Database
}

func NewTodoService(DB *bootstraps.Database) *TodoService {
	return &TodoService{DB}
}

// func Create(todo *model.Todo) error {
// 	return bootstraps.DB.Create(&todo).Error
// }

// func Update(id string, todo map[string]interface{}) error {
// 	return bootstraps.DB.Model(&model.Todo{}).Where("id = ?", id).Updates(todo).Error
// }

// func (db TodoService) GetAll() (map[string]interface{}, error) {
// 	var todos []model.Todo

// 	err := db.Find(&todos).Error
// 	if err != nil {
// 		fmt.Println("Error while getting")
// 		return nil, err
// 	}

// 	return gin.H{"data": todos}, nil
// }

// get todo for a specific user
func (db TodoService) GetUserTodo() (map[string]interface{}, error) {
	var todos []model.Todo

	// err := db.Where("user_id = ?", userId).Find(&todos).Error
	err := db.Find(&todos).Error
	if err != nil {
		fmt.Println("Error while getting")
		return nil, err
	}

	return gin.H{"data": todos}, nil
}

// func GetById(id string) (todo *model.Todo, err error) {
// 	return todo, bootstraps.DB.First(&todo, "id = ?", id).Error
// }
