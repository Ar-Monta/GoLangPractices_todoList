package commands

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ArMo-Team/GoLangPractices_todoList/internal/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

type MigrateCommand struct{}

func (c *MigrateCommand) Execute() {
	// Connect to the database
	db := database.Connect()

	// Run the migrations
	err := MigrateDB(db)
	if err != nil {
		return
	}

	// Close the database
	closeErr := db.Close()
	if closeErr != nil {
		return
	}

	fmt.Println("All migrations completed successfully.")
}

func MigrateDB(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal("Failed to initialize database driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/migrations",
		"mysql", driver)
	if err != nil {
		log.Fatal("Failed to initialize migration:", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("Failed to apply migrations:", err)
		return err
	}

	return nil
}
