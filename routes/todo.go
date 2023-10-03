package routes

import (
	"todo/controller"
	"todo/infrastructure"
	"todo/middleware"
)

func SetupRoutes(handler *infrastructure.Handler, todoController *controller.TodoController, userController *controller.UserController, middleware *middleware.AuthMiddleware) {
	// handler.Gin.GET("/all-todo", todoController.GetAllTodo)

	v1 := handler.Gin.Group("/")
	{
		v1.POST("/login", userController.Login)
	}

	v2 := handler.Gin.Group("/todo").Use(middleware.ValidateToken())
	{
		v2.GET("", todoController.GetAllTodo)
	}
}
