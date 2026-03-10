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
	"github.com/AaravShetty15/go-todo-app/config"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"context"
	
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	cfg := config.LoadConfig()

	fmt.Println("Starting Todo API server...")

	// Connect to SQLite database
	db, err := sql.Open("sqlite3", cfg.DBPath)
	if err != nil {
		log.Fatal(err)
	}

	// Redis client
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{ Addr: "localhost:6379",
	})

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
	service := services.NewTodoService(repo, rdb, ctx)
	handler := handlers.NewTodoHandler(service)

	// Setup routes
	router := routes.SetupRoutes(handler)

	// Start server
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}