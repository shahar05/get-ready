package tests

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"phonebook-api/contacts"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	// Initialize test database
	var err error
	connStr := "host=localhost port=5432 user=postgres password=pass1234 dbname=contactdb sslmode=disable"
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
