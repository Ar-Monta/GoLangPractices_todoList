package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/ArMo-Team/GoLangPractices_todoList/internal/app/todo"
	"github.com/ArMo-Team/GoLangPractices_todoList/internal/infrastructure/persistence"
	"github.com/ArMo-Team/GoLangPractices_todoList/migrations"
	"github.com/ArMo-Team/GoLangPractices_todoList/pkg/api"
)

func main() {
	// Load variables from .env file into environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	// Access environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	// You can add password here by using /* dbPassword := os.Getenv("DB_PASSWORD") */

	// Open connection to database
	db, err := sql.Open("mysql", dbUser+"@tcp("+dbHost+":"+dbPort+")/"+os.Getenv("DB_NAME"))
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
