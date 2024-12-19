package todo

import (
	"boilerplate-api/lib/config"
	"errors"

	"gorm.io/gorm"
)

// type TodoRepository interface {
// 	CreateTodo(todo *TodoModel) error
// 	GetAllTodos() ([]TodoModel, error)
// 	UpdateTodo(todo *TodoModel) error
// 	DeleteTodo(id uint) error
// }

type TodoRepository struct {
	db     *config.Database
	logger config.Logger
}

func NewTodoRepository(db *config.Database, logger config.Logger) TodoRepository {
	return TodoRepository{
		db:     db,
		logger: logger,
	}
}

func (r *TodoRepository) CreateTodo(todo *TodoModel) error {
	return r.db.Create(todo).Error
}

func (r *TodoRepository) GetAllTodos() ([]TodoModel, error) {
	var todos []TodoModel
	err := r.db.Order("id DESC").Find(&todos).Error
	return todos, err
}

func (r *TodoRepository) UpdateTodo(todo *TodoModel) error {
	return r.db.Save(todo).Error
}

func (r *TodoRepository) DeleteTodo(id uint) error {
	return r.db.Delete(&TodoModel{}, id).Error
}

func (r *TodoRepository) FindTodoByID(id uint) (*TodoModel, error) {
	var todo TodoModel
	err := r.db.First(&todo, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &todo, err
}
