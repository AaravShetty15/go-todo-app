package routes

import (
	"github.com/gorilla/mux"
	"github.com/AaravShetty15/go-todo-app/handlers"
)

func SetupRoutes(todoHandler *handlers.TodoHandler) *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	router.HandleFunc("/todos", todoHandler.GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", todoHandler.GetTodoByID).Methods("GET")
	router.HandleFunc("/todos/{id}", todoHandler.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", todoHandler.DeleteTodo).Methods("DELETE")

	return router
}