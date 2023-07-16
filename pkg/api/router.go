package api

import (
	"net/http"

	"github.com/ArMo-Team/GoLangPractices_todoList/internal/app/todo"
	"github.com/gorilla/mux"
)

type Router struct {
	TodoHandler *todo.TodoHandler
}

func (r *Router) SetupRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/todos", r.TodoHandler.GetTodosHandler).Methods(http.MethodGet)
	router.HandleFunc("/todos/{id}", r.TodoHandler.GetTodoByIDHandler).Methods(http.MethodGet)
	router.HandleFunc("/todos", r.TodoHandler.CreateTodoHandler).Methods(http.MethodPost)
	router.HandleFunc("/todos/{id}", r.TodoHandler.UpdateTodoCompletedHandler).Methods(http.MethodPut)
	router.HandleFunc("/todos/{id}", r.TodoHandler.DeleteTodoHandler).Methods(http.MethodDelete)
	// Define routes for UpdateTodoHandler

	return router
}
