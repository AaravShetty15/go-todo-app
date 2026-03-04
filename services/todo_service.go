package services

import (
	"github.com/AaravShetty15/go-todo-app/models"
	"github.com/AaravShetty15/go-todo-app/repository"
)

type TodoService struct {
	Repo *repository.TodoRepository
}

// constructor
func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{
		Repo: repo,
	}
}

// create Todo
func (s *TodoService) CreateTodo(todo models.Todo) error {
	return s.Repo.CreateTodo(todo)
}

// get All Todos
func (s *TodoService) GetTodos() ([]models.Todo, error) {
	return s.Repo.GetTodos()
}

// get Todo by ID
func (s *TodoService) GetTodoByID(id int) (models.Todo, error) {
	return s.Repo.GetTodoByID(id)
}

// update Todo
func (s *TodoService) UpdateTodo(todo models.Todo) error {
	return s.Repo.UpdateTodo(todo)
}

// delete Todo
func (s *TodoService) DeleteTodo(id int) error {
	return s.Repo.DeleteTodo(id)
}