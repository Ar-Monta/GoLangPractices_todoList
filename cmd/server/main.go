package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ArMo-Team/GoLangPractices_todoList/internal/app/todo"
	"github.com/ArMo-Team/GoLangPractices_todoList/internal/infrastructure/persistence"
	"github.com/ArMo-Team/GoLangPractices_todoList/migrations"
	"github.com/ArMo-Team/GoLangPractices_todoList/pkg/api"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/todo-app-db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	migrations.MigrateDB(db)

	todoRepository := persistence.NewTodoRepository(db)
	todoService := todo.NewTodoService(todoRepository)
	todoHandler := &todo.TodoHandler{Service: todoService}
	router := api.Router{TodoHandler: todoHandler}
	http.Handle("/", router.SetupRoutes())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
