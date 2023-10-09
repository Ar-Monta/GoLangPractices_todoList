package api

import (
	"github.com/ArMo-Team/GoLangPractices_todoList/internal/app/user"
	"net/http"

	"github.com/ArMo-Team/GoLangPractices_todoList/internal/app/todo"
	"github.com/gorilla/mux"
)

type Router struct {
	TodoHandler *todo.TodoHandler
	UserHandler *user.UserHandler
}

func (r *Router) SetupRoutes() http.Handler {
	router := mux.NewRouter()

	// Todos end-points
	router.HandleFunc("/todos", r.TodoHandler.GetTodosHandler).Methods(http.MethodGet)
	router.HandleFunc("/todos/{id}", r.TodoHandler.GetTodoByIDHandler).Methods(http.MethodGet)
	router.HandleFunc("/todos", r.TodoHandler.CreateTodoHandler).Methods(http.MethodPost)
	router.HandleFunc("/todos/{id}", r.TodoHandler.UpdateTodoCompletedHandler).Methods(http.MethodPut)
	router.HandleFunc("/todos/{id}", r.TodoHandler.DeleteTodoHandler).Methods(http.MethodDelete)
	// Define routes for UpdateTodoHandler

	// Users end-points
	router.HandleFunc("/users", r.UserHandler.GetUsersHandler).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", r.UserHandler.GetUsersHandler).Methods(http.MethodGet)
	router.HandleFunc("/users", r.UserHandler.CreateUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", r.UserHandler.UpdateUserVerifiedAtHandler).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", r.UserHandler.DeleteUserHandler).Methods(http.MethodDelete)

	return router
}
