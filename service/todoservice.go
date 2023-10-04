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

func (db TodoService) CreateTodo(todo *model.Todo) error {
	return db.Create(&todo).Error

}

func (db TodoService) Update(id string, todo map[string]interface{}) error {
	fmt.Println(todo)
	return db.Model(&model.Todo{}).Where("id = ?", id).Updates(todo).Error
}

func (db TodoService) GetAll() (map[string]interface{}, error) {
	var todos []model.Todo

	err := db.Find(&todos).Error
	if err != nil {
		fmt.Println("Error while getting")
		return nil, err
	}

	return gin.H{"data": todos}, nil
}

// get todo for a specific user
func (db TodoService) GetUserTodo(userId string) (map[string]interface{}, error) {
	var todos []model.Todo

	err := db.Where("user_id = ?", userId).Find(&todos).Error
	if err != nil {
		fmt.Println("Error while getting")
		return nil, err
	}

	return gin.H{"data": todos}, nil
}

func (db TodoService) GetById(id string) (todo *model.Todo, err error) {
	return todo, db.First(&todo, "id = ?", id).Error
}
