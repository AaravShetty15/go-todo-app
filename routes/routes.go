package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/AaravShetty15/go-todo-app/handlers"
	"github.com/AaravShetty15/go-todo-app/middleware"
)

func SetupRoutes(todoHandler *handlers.TodoHandler) *mux.Router {

	router := mux.NewRouter()
	router.Use(middleware.Logging)

	router.Handle("/todos", middleware.BasicAuth(http.HandlerFunc(todoHandler.CreateTodo)),).Methods("POST")
	router.HandleFunc("/todos", todoHandler.GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", todoHandler.GetTodoByID).Methods("GET")
	router.Handle("/todos/{id}", middleware.BasicAuth(http.HandlerFunc(todoHandler.UpdateTodo)),).Methods("PUT")
	router.Handle("/todos/{id}",
	middleware.BasicAuth(http.HandlerFunc(todoHandler.DeleteTodo)),).Methods("DELETE")
	router.HandleFunc("/suggest-task", todoHandler.SuggestTask).Methods("GET")
	router.HandleFunc("/weather", todoHandler.GetWeather).Methods("GET")

	return router
}