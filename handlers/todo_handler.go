package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/AaravShetty15/go-todo-app/utils"
	"github.com/AaravShetty15/go-todo-app/models"
	"github.com/AaravShetty15/go-todo-app/services"
	"github.com/AaravShetty15/go-todo-app/external"
)

type TodoHandler struct {
	Service *services.TodoService
}

// constructor
func NewTodoHandler(service *services.TodoService) *TodoHandler {
	return &TodoHandler{
		Service: service,
	}
}

// POST /todos handles CreateTodo()
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {

	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Service.CreateTodo(todo)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
}

	utils.WriteJSON(w, http.StatusCreated, utils.APIResponse{
	Success: true,
	Message: "Todo created successfully",
	Data:    todo,
	})
}

// GET /todos
func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 5

	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err == nil && p > 0 {
			page = p
		}
	}

	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err == nil && l > 0 {
			limit = l
		}
	}

	offset := (page - 1) * limit

	todos, err := h.Service.GetTodosPaginated(limit, offset)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.APIResponse{
		Success: true,
		Data:    todos,
	})
}

// GET /todos/{id}
func (h *TodoHandler) GetTodoByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, err := h.Service.GetTodoByID(id)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.APIResponse{
	Success: true,
	Data:    todo,
	})
}

// PUT /todos/{id}
func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var todo models.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo.ID = id

	err = h.Service.UpdateTodo(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.APIResponse{
	Success: true,
	Message: "Todo updated successfully",
	Data:    todo,
	})
}

// DELETE /todos/{id}
func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.APIResponse{
	Success: true,
	Message: "Todo deleted successfully",
	})
}

func (h *TodoHandler) SuggestTask(w http.ResponseWriter, r *http.Request) {

	task, err := external.GetSuggestedTask()
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.APIResponse{
		Success: true,
		Data:    task,
	})
}

func (h *TodoHandler) GetWeather(w http.ResponseWriter, r *http.Request) {

	// Bangalore coordinates example
	weather, err := external.GetWeather(12.9716, 77.5946)

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.APIResponse{
		Success: true,
		Data:    weather,
	})
}