package todo

import (
	"boilerplate-api/lib/config"
	"boilerplate-api/lib/router"
)

func SetupRoutes(
	logger config.Logger,
	router router.Router,
	todoController TodoController,

) {
	logger.Info(" Setting up user routes")
	todos := router.V1.Group("/todos")

	todos.GET("/", todoController.GetAllTodosHandler)
	todos.GET("/:id", todoController.GetTodoByIdHandler)
	todos.POST("/", todoController.CreateTodoHandler)
	todos.PUT("/:id", todoController.UpdateTodoHandler)
	todos.DELETE("/:id", todoController.DeleteTodoHandler)
}
