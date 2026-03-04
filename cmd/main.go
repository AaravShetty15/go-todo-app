package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"github.com/AaravShetty15/go-todo-app/repository"
	"github.com/AaravShetty15/go-todo-app/services"
	"github.com/AaravShetty15/go-todo-app/handlers"
	"github.com/AaravShetty15/go-todo-app/routes"
	
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	fmt.Println("Starting Todo API server...")

	// Connect to SQLite database
	db, err := sql.Open("sqlite3", "./todos.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create todos table if it doesn't exist
	createTable := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		completed BOOLEAN,
		created_at DATETIME,
		updated_at DATETIME
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize layers
	repo := repository.NewTodoRepository(db)
	service := services.NewTodoService(repo)
	handler := handlers.NewTodoHandler(service)

	// Setup routes
	router := routes.SetupRoutes(handler)

	// Start server
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}