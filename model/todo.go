package model

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	UserID      string `json:"user_id" form:"user_id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description"`
}

type UpdateTodo struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
}
