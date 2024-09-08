package main

import (
	"log"
	"net/http"
	"phonebook-api/contacts"
	"phonebook-api/database"
	"phonebook-api/middleware"
	"phonebook-api/token"
	"phonebook-api/users"

	"github.com/alicebob/miniredis/server"
	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

func main() {
	db := database.Init()

	r := mux.NewRouter()

	// Create PASETO token maker
	pasetoMaker, err := token.NewPasetoMaker("your-secret-key") // TODO: Replace with your actual key
	if err != nil {
		log.Fatalf("Failed to create token maker: %v", err)
	}

	server := &server.Server{
		db:         db,
		tokenMaker: pasetoMaker,
	}

	// Register the HealthCheckHandler
	r.HandleFunc("/", HealthCheckHandler).Methods("GET")

	// Register Contact Handlers with Auth Middleware
	r.Use(middleware.AuthMiddleware(pasetoMaker))

	users.RegisterRoutes(r, server)
	contacts.RegisterRoutes(r, server)

	// Start server
	portServer := "8080"
	log.Printf("Server is running on port %s", portServer)
	log.Fatal(http.ListenAndServe(":"+portServer, r))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Health Check ok"))
}
