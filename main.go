package main

import (
	"log"
	"net/http"
	"phonebook-api/contacts"
	"phonebook-api/database"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

func main() {
	db := database.Init()
	service := contacts.NewService(db)
	server := contacts.NewServer(db, service)

	r := mux.NewRouter()
	server.RegisterRoutes(r)

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

// func main() {

// 	db := database.Init()

// 	r := mux.NewRouter()

// 	// Register the HealthCheckHandler
// 	r.HandleFunc("/", HealthCheckHandler).Methods("GET")

// 	// Register Contact Handlers
// 	contacts.RegisterRoutes(r, db)

// 	// Start server
// 	portServer := "8080"
// 	log.Printf("Server is running on port %s", portServer)
// 	log.Fatal(http.ListenAndServe(":"+portServer, r))
// }

// func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Health Check ok"))
// }
