package main

import (
	// "todo/bootstraps"

	"todo/serve"

	"go.uber.org/fx"
)

// func init() {
// 	bootstraps.LoadEnv()
// 	bootstraps.ConnectToDB()
// }

func main() {
	// r := gin.Default()

	// fmt.Println("API connected")

	// v1 := r.Group("/todo").Use(middleware.ValidateToken())
	// {
	// 	//todo routes
	// 	v1.POST("", controller.CreateTodo)
	// 	v1.GET("/alltodo", controller.GetAllTodo)
	// 	v1.PATCH("/:id", controller.UpdateTodo)
	// 	v1.GET("/:id", controller.GetTodoById)
	// }

	// v2 := r.Group("/")
	// {
	// 	//authentication routes
	// 	v2.POST("signup", controller.Signup)
	// 	v2.POST("login", controller.Login)
	// }

	// r.GET("/health-check", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Todo api up and running",
	// 	})
	// })

	// portStr := os.Getenv("SERVER_PORT")
	// port, err := strconv.Atoi(portStr)

	// if err != nil {
	// 	fmt.Println("Error while loading port from env", err)
	// }

	fx.New(
		serve.Module,
	).Run()
}
