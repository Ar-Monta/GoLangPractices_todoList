package commands

import (
	"database/sql"
	"fmt"
	"github.com/ArMo-Team/GoLangPractices_todoList/internal/app/todo"
	"github.com/ArMo-Team/GoLangPractices_todoList/internal/app/user"
	"github.com/ArMo-Team/GoLangPractices_todoList/internal/database"
	"github.com/ArMo-Team/GoLangPractices_todoList/internal/infrastructure/persistence"
	"github.com/ArMo-Team/GoLangPractices_todoList/pkg/api"
	"log"
	"net/http"
)

type StartCommand struct{}

func (c *StartCommand) Execute() {
	err := database.WithDB(
		start,
	)

	if err != nil {
		log.Fatal("failed to start application:", err)
	}

	fmt.Println("Stopped listening on port 8080")
}

func start(db *sql.DB) error {

	router := api.Router{
		TodoHandler: createTodoHandler(db),
		UserHandler: createUserHandler(db),
	}
	http.Handle("/", router.SetupRoutes())

	fmt.Println("Started listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	return nil
}

func createTodoHandler(db *sql.DB) *todo.TodoHandler {
	todoRepository := persistence.NewTodoRepository(db)
	todoService := todo.NewTodoService(todoRepository)
	todoHandler := &todo.TodoHandler{Service: todoService}
	return todoHandler
}

func createUserHandler(db *sql.DB) *user.UserHandler {
	userRepository := persistence.NewUserRepository(db)
	fmt.Println("Setting up handlers")
	userService := user.NewUserService(userRepository)
	userHandler := &user.UserHandler{Service: userService}
	return userHandler
}
