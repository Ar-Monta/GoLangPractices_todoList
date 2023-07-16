package todo

import (
	"encoding/json"
	"github.com/ArMo-Team/GoLangPractices_todoList/internal/app/domain"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type TodoHandler struct {
	Service TodoService
}

func (h *TodoHandler) GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := h.Service.GetTodos()
	if err != nil {
		// Handle error and return HTTP response
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todosJSON, err := json.Marshal(todos)
	if err != nil {
		// Handle JSON conversion error and return HTTP response
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set HTTP headers
	w.Header().Set("Content-Type", "application/json")

	// Write JSON response
	_, err = w.Write(todosJSON)

	if err != nil {
		// Handle Write JSON response error and return HTTP response
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *TodoHandler) GetTodoByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["id"]

	todo, err := h.Service.GetTodoByID(todoID)
	if err != nil {
		// Handle error and return HTTP response
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if todo == nil {
		// Handle ID not found and return HTTP response
		http.Error(w, "No todo with this ID found!", http.StatusNotFound)
	}

	todoJSON, err := json.Marshal(todo)
	if err != nil {
		// Handle JSON conversion error and return HTTP response
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set HTTP headers
	w.Header().Set("Content-Type", "application/json")

	// Write JSON response
	_, err = w.Write(todoJSON)

	if err != nil {
		// Handle Write JSON response error and return HTTP response
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *TodoHandler) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	// Parse the request body into a map
	var requestData map[string]interface{}
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new todo with title and description
	todo := domain.Todo{
		Title:       requestData["title"].(string),
		Description: requestData["description"].(string),
		// Other fields will be handled by default values in the database
	}

	// Create the todo using the service
	err = h.Service.CreateTodo(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the created todo to JSON
	todoJSON, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the appropriate HTTP headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Write the JSON response
	w.Write(todoJSON)
}

func (h *TodoHandler) UpdateTodoCompletedHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the todo ID from the request URL parameters
	vars := mux.Vars(r)
	todoID := vars["id"]

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	// Parse the request body into a map
	var requestData map[string]interface{}
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the completed status from the request body
	completed, ok := requestData["completed"].(bool)
	if !ok {
		http.Error(w, "Invalid completed status", http.StatusBadRequest)
		return
	}

	// Update the todo completed status using the service
	err = h.Service.UpdateTodoCompleted(todoID, completed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}

func (h *TodoHandler) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["id"]

	err := h.Service.DeleteTodo(todoID)
	if err != nil {
		// Handle error and return HTTP response
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Set HTTP headers
	w.Header().Set("Content-Type", "application/json")
}
