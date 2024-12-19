package todo

type TodoService struct {
	repository TodoRepository
}

func NewTodoService(repository TodoRepository) TodoService {
	return TodoService{
		repository: repository,
	}
}

func (s *TodoService) CreateTodo(todo *TodoModel) error {
	return s.repository.CreateTodo(todo)
}

func (s *TodoService) GetAllTodos() ([]TodoModel, error) {
	return s.repository.GetAllTodos()
}

func (s *TodoService) UpdateTodo(todo *TodoModel) error {
	return s.repository.UpdateTodo(todo)
}

func (s *TodoService) DeleteTodo(id uint) error {
	return s.repository.DeleteTodo(id)
}

func (s *TodoService) FindTodoByID(id uint) (*TodoModel, error) {
	return s.repository.FindTodoByID(id)
}
