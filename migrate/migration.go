package main

import (
	"fmt"
	"todo/bootstraps"
	"todo/model"
)

func init() {
	bootstraps.LoadEnv()
	bootstraps.ConnectToDB()
}

func main() {
	fmt.Println("Migration Started")
	bootstraps.DB.AutoMigrate(&model.Todo{}, &model.User{})
}
