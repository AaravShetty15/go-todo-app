package services

import (
	"context"
	"encoding/json"
	"errors"
	"time"
	"github.com/redis/go-redis/v9"
	"github.com/AaravShetty15/go-todo-app/models"
	"github.com/AaravShetty15/go-todo-app/repository"
)

type TodoService struct {
	Repo *repository.TodoRepository
	RDB  *redis.Client
	Ctx  context.Context
}

// constructor
func NewTodoService(repo *repository.TodoRepository, rdb *redis.Client, ctx context.Context) *TodoService {
	return &TodoService{
		Repo: repo,
		RDB:  rdb,
		Ctx:  ctx,
	}
}

// create Todo
func (s *TodoService) CreateTodo(todo models.Todo) error {

	if todo.Title == "" {
		return errors.New("title cannot be empty")
	}

	if len(todo.Description) < 5 {
		return errors.New("description must be at least 5 characters")
	}

	err := s.Repo.CreateTodo(todo)

	if err == nil {
		s.RDB.Del(s.Ctx, "todos") // clear cache
	}

	return err
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

	if todo.Title == "" {
		return errors.New("title cannot be empty")
	}

	if len(todo.Description) < 5 {
		return errors.New("description must be at least 5 characters")
	}

	err := s.Repo.UpdateTodo(todo)

	if err == nil {
		s.RDB.Del(s.Ctx, "todos") // clear cache
	}

	return err
}

// delete Todo
func (s *TodoService) DeleteTodo(id int) error {

	err := s.Repo.DeleteTodo(id)

	if err == nil {
		s.RDB.Del(s.Ctx, "todos") // clear cache
	}

	return err
}

func (s *TodoService) GetTodosPaginated(limit, offset int) ([]models.Todo, error) {

	cacheKey := "todos"

	val, err := s.RDB.Get(s.Ctx, cacheKey).Result()

	if err == nil {

		var todos []models.Todo
		json.Unmarshal([]byte(val), &todos)

		return todos, nil
	}

	todos, err := s.Repo.GetTodosPaginated(limit, offset)

	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(todos)

	s.RDB.Set(s.Ctx, cacheKey, data, 5*time.Minute)

	return todos, nil
}