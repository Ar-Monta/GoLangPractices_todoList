package main

import (
	"database/sql"
	"fmt"
	"github.com/ArMo-Team/GoLangPractices_todoList/database"
	"github.com/ArMo-Team/GoLangPractices_todoList/migrations"
	"log"
)

func main() {
	err := database.WithDB(
		func(db *sql.DB) error {
			fmt.Println("Migrating to database")
			return migrations.MigrateDB(db)
		},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
}
