package routes

import (
	"todo/controller"
	"todo/infrastructure"
	"todo/middleware"
)

func SetupTodoRoutes(handler *infrastructure.Handler, todoController *controller.TodoController, userController *controller.UserController, middleware *middleware.AuthMiddleware) {
	// handler.Gin.GET("/all-todo", todoController.GetAllTodo)

	todo := handler.Gin.Group("/todo").Use(middleware.ValidateToken())
	{
		todo.POST("", todoController.CreateTodo)
		todo.GET("/alltodo", todoController.GetAllTodo)
		todo.PATCH("/:id", todoController.UpdateTodo)
		todo.GET("/:id", todoController.GetTodoById)
	}
}
