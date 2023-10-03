package controller

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewTodoController),
	fx.Provide(NewUserController),
)
