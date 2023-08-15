package main

import (
	"database/sql"
	"fmt"
	"github.com/ArMo-Team/GoLangPractices_todoList/internal/app/todo"
	"github.com/ArMo-Team/GoLangPractices_todoList/internal/commands"
	"github.com/ArMo-Team/GoLangPractices_todoList/internal/infrastructure/persistence"
	"github.com/ArMo-Team/GoLangPractices_todoList/pkg/api"
	"log"
	"net/http"
	"os"
)

func main() {
	// Parse the command and its arguments
	command, err := commands.ParseCommand()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	// Execute the chosen command
	command.Execute()

	//err := database.WithDB(
	//	start,
	//)
	//
	//if err != nil {
	//	log.Fatal("failed to start application:", err)
	//}
}

func start(db *sql.DB) error {
	todoRepository := persistence.NewTodoRepository(db)
	fmt.Println("Setting up handlers")
	todoService := todo.NewTodoService(todoRepository)
	todoHandler := &todo.TodoHandler{Service: todoService}
	router := api.Router{TodoHandler: todoHandler}
	http.Handle("/", router.SetupRoutes())

	fmt.Println("Started listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	return nil
}
