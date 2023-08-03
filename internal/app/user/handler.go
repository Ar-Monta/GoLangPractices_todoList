package user

import (
	"encoding/json"
	"github.com/ArMo-Team/GoLangPractices_todoList/internal/app/domain"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	Service UserService
}

func (h *UserHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetUsers()
	if err != nil {
		// Handle error and return HTTP response
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	usersJSON, err := json.Marshal(users)
	if err != nil {
		// Handle JSON conversion error and return HTTP response
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set HTTP headers
	w.Header().Set("Content-Type", "application/json")

	// Write JSON response
	_, err = w.Write(usersJSON)

	if err != nil {
		// Handle Write JSON response error and return HTTP response
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *UserHandler) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	user, err := h.Service.GetUserByID(userID)
	if err != nil {
		// Handle error and return HTTP response
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if user == nil {
		// Handle ID not found and return HTTP response
		http.Error(w, "No user with this ID found!", http.StatusNotFound)
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		// Handle JSON conversion error and return HTTP response
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set HTTP headers
	w.Header().Set("Content-Type", "application/json")

	// Write JSON response
	_, err = w.Write(userJSON)

	if err != nil {
		// Handle Write JSON response error and return HTTP response
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
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

	// Create a new user with title and description
	user := domain.User{
		FirstName: requestData["first_name"].(string),
		LastName:  requestData["last_name"].(string),
		// Other fields will be handled by default values in the database
	}

	// Create the user using the service
	err = h.Service.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the created user to JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the appropriate HTTP headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Write the JSON response
	_, err = w.Write(userJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) UpdateUserVerifiedAtHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the request URL parameters
	vars := mux.Vars(r)
	userID := vars["id"]

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

	// Get the verified_at status from the request body
	verifiedAt, ok := requestData["verified_at"].(time.Time)
	if !ok {
		http.Error(w, "Invalid verified_at status", http.StatusBadRequest)
		return
	}

	// Update the user verified_at status using the service
	err = h.Service.UpdateUserVerifiedAt(userID, verifiedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	err := h.Service.DeleteUser(userID)
	if err != nil {
		// Handle error and return HTTP response
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Set HTTP headers
	w.Header().Set("Content-Type", "application/json")
}
