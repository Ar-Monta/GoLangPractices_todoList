package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Connect() *sql.DB {
	// Load variables from .env file into environment variables
	fmt.Println("Reading .env file")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Access environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	// You can add password here by using /* dbPassword := os.Getenv("DB_PASSWORD") */

	// Open connection to database
	fmt.Println("Opening connection to database")
	db, err := sql.Open("mysql", dbUser+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Failed to close database connection:", err)
		}
	}(db)

	fmt.Println("Pinging connection to database")
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}
	return db
}

func WithDB(tasks ...func(*sql.DB) error) error {
	// Load variables from .env file into environment variables
	fmt.Println("Reading .env file")
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	// Access environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	// You can add password here by using /* dbPassword := os.Getenv("DB_PASSWORD") */

	// Open connection to database
	fmt.Println("Opening connection to database")
	db, err := sql.Open("mysql", dbUser+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Ping the database to ensure the connection is established
	fmt.Println("Pinging connection to database")
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping the database: %w", err)
	}

	// Execute the provided tasks using the database connection
	for _, task := range tasks {
		err = task(db)
		if err != nil {
			// Don't close the database connection immediately; continue with other tasks
			fmt.Println("Task failed:", err)
		}
	}

	// Close the database connection
	closeErr := db.Close()
	if closeErr != nil {
		return fmt.Errorf("failed to close database connection: %w", closeErr)
	}

	return nil
}
