package tests

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"phonebook-api/contacts"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	// Load environment variables from .env file
	err = godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize test database
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	testDB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Ensure the database connection works
	if err := testDB.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	// Set the database connection for your application
	contacts.SetDB(testDB)

	// Run tests
	exitCode := m.Run()

	// Clean up
	testDB.Close()

	os.Exit(exitCode)
}
