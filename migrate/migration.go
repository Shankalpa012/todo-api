package main

import (
	"fmt"
	"todo/bootstraps"
	"todo/model"
)

func init() {
	bootstraps.LoadEnv()
	bootstraps.NewDatabase()
}

func main() {
	fmt.Println("Migration Started")
	bootstraps.NewDatabase().AutoMigrate(&model.Todo{}, &model.User{})
}
