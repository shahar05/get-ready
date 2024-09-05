package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"phonebook-api/contacts"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

func main() {

	db := InitializeDB()

	initTable(db)

	r := mux.NewRouter()

	// Register Contact Handlers
	contacts.RegisterRoutes(r, db)

	// Start server
	portServer := "8080"

	log.Printf("Server is running on port %s", portServer)
	log.Fatal(http.ListenAndServe(":"+portServer, r))
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

// InitializeDB initializes the database connection
func InitializeDB() *sql.DB {
	var err error

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	/// I must be removed as I am a comment
	// fmt.Printf("This is the ConnSTR: %s", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	return db
}
