package repository

import (
	"database/sql"
	"time"
	"github.com/AaravShetty15/go-todo-app/models"
)

// holds the db connection
type TodoRepository struct {
	DB *sql.DB
}

// constructor
func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		DB: db,
	}
}

// create a new todo
func (r *TodoRepository) CreateTodo(todo models.Todo) error {
	query := `
	INSERT INTO todos (title, description, completed, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?)
	`

	_, err := r.DB.Exec(
		query,
		todo.Title,
		todo.Description,
		todo.Completed,
		time.Now(),
		time.Now(),
	)

	return err
}

// get all todos
func (r *TodoRepository) GetTodos() ([]models.Todo, error) {

	rows, err := r.DB.Query(`SELECT id, title, description, completed, created_at, updated_at FROM todos`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo

		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

// get todo by ID
func (r *TodoRepository) GetTodoByID(id int) (models.Todo, error) {

	var todo models.Todo

	query := `SELECT id, title, description, completed, created_at, updated_at FROM todos WHERE id = ?`

	err := r.DB.QueryRow(query, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)

	return todo, err
}

// update a todo
func (r *TodoRepository) UpdateTodo(todo models.Todo) error {

	query := `
	UPDATE todos
	SET title = ?, description = ?, completed = ?, updated_at = ?
	WHERE id = ?
	`

	_, err := r.DB.Exec(
		query,
		todo.Title,
		todo.Description,
		todo.Completed,
		time.Now(),
		todo.ID,
	)

	return err
}

// delete todo
func (r *TodoRepository) DeleteTodo(id int) error {

	query := `DELETE FROM todos WHERE id = ?`

	_, err := r.DB.Exec(query, id)

	return err
}	