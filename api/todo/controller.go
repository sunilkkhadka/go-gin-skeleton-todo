package todo

import (
	"boilerplate-api/lib/config"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	logger      config.Logger
	todoService TodoService
	env         config.Env
}

func NewTodoController(
	logger config.Logger,
	todoService TodoService,
	env config.Env,
) TodoController {
	return TodoController{
		logger:      logger,
		todoService: todoService,
		env:         env,
	}
}
func (h *TodoController) CreateTodoHandler(c *gin.Context) {
	var todo TodoModel
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
		return
	}

	if err := h.todoService.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create a todo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Todo Created Successfully",
	})
}

func (h *TodoController) GetAllTodosHandler(c *gin.Context) {
	todos, err := h.todoService.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (h *TodoController) UpdateTodoHandler(c *gin.Context) {
	var todo TodoEditRequest
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
		return
	}

	var existingTodo *TodoModel
	existingTodo, err := h.todoService.FindTodoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Todo with this id not found"})
		return
	}

	if existingTodo == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	newTodo := &TodoModel{
		ID:          uint(id),
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
		CreatedAt:   existingTodo.CreatedAt,
		UpdatedAt:   time.Now(),
		DeletedAt:   existingTodo.DeletedAt,
	}

	if err := h.todoService.UpdateTodo(newTodo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot update todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo Updated Successfully",
	})
}

func (h *TodoController) DeleteTodoHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var existingTodo *TodoModel
	existingTodo, err := h.todoService.FindTodoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Todo with this id not found"})
		return
	}

	if existingTodo == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo item not found"})
		return
	}

	if err := h.todoService.DeleteTodo(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted Successfully"})
}

func (h *TodoController) GetTodoByIdHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.todoService.FindTodoByID(uint(id))
	if err != nil || todo == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}
