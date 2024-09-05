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

	r := mux.NewRouter()

	// Register Contact Handlers
	contacts.RegisterRoutes(r, db)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

// InitializeDB initializes the database connection
func InitializeDB() *sql.DB {
	var err error

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	return db
}
