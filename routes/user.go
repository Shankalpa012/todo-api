package routes

import (
	"todo/controller"
	"todo/infrastructure"
	"todo/middleware"
)

func SetupUserRoutes(handler *infrastructure.Handler, todoController *controller.TodoController, userController *controller.UserController, middleware *middleware.AuthMiddleware) {
	// handler.Gin.GET("/all-todo", todoController.GetAllTodo)

	user := handler.Gin.Group("/")
	{
		user.POST("login", userController.Login)
		user.POST("signup", userController.Signup)
	}
}
