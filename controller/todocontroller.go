package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"todo/model"
	"todo/service"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoService *service.TodoService
}

func NewTodoController(todoService *service.TodoService) *TodoController {
	return &TodoController{
		todoService: todoService,
	}
}

func (t TodoController) CreateTodo(c *gin.Context) {
	data := model.Todo{}

	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println("Error While binding the JSON")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error While binding the JSON",
		})
		return
	}

	err := t.todoService.CreateTodo(&model.Todo{
		Title:       data.Title,
		Description: data.Description,
		UserID:      data.UserID,
	})

	if err != nil {
		fmt.Println("Error while adding to DB")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "TODO Created successfully",
	})

}

func (t TodoController) UpdateTodo(ctx *gin.Context) {
	uuid := ctx.Param("id")

	fmt.Println(uuid)

	var todo model.UpdateTodo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error While parsing"})
		return
	}

	finalData := structs.Map(todo)

	if err := t.todoService.Update(uuid, finalData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error While updating"})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "TODO Updated successfully",
	})

}

func (tc TodoController) GetAllTodo(c *gin.Context) {

	userId := c.MustGet("id").(float64)

	todos, err := tc.todoService.GetUserTodo(strconv.Itoa(int(userId)))
	if err != nil {
		fmt.Println("Error while fetching all the data", err)
		return
	}
	if err != nil {
		fmt.Println("Error while fetching all the data", err)
		return
	}

	c.JSON(200, gin.H{
		"data": todos,
	})
}

func (t TodoController) GetTodoById(c *gin.Context) {
	id := c.Param("id")

	todo, err := t.todoService.GetById(id)
	if err != nil {
		fmt.Println("Error while fetching single data", err)
		return
	}
	fmt.Println(todo)
	c.JSON(200, gin.H{
		"data": todo,
	})
}
