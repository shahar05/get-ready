package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// Init initializes the database connection
func Init() *sql.DB {
	var err error

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	// Establish DB connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	// Ping DB server
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	// Initialize the contacts table
	if err := initTable(db); err != nil {
		log.Fatalf("failed to initialize the contacts table: %v", err)
	}

	return db
}

func initTable(db *sql.DB) error {
	// Define the SQL statement for creating the table
	query := `
    CREATE TABLE IF NOT EXISTS contacts (
        id SERIAL PRIMARY KEY,
        first_name VARCHAR(100),
        last_name VARCHAR(100),
        phone VARCHAR(20),
        address TEXT
    );`

	// Execute the SQL statement
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating contacts table: %w", err)
	}
	return nil
}
