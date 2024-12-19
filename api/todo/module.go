package todo

import (
	"go.uber.org/fx"
)

var Module = fx.Module("todo",
	fx.Options(
		fx.Provide(
			NewTodoRepository,
			NewTodoService,
			NewTodoController),
		fx.Invoke(SetupRoutes),
	))
