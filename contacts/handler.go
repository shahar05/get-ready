package contacts

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"phonebook-api/utils"
	"strconv"

	"github.com/gorilla/mux"
)

// Service struct that implements the ContactService interface
type Service struct {
	db *sql.DB
}

// NewService creates a new service instance
func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

// ContactService interface defines methods to interact with contacts
type ContactService interface {
	GetContacts(limit, offset int) ([]Contact, error)
	AddContact(contact Contact) (string, error)
}

// Server struct contains the DB instance and the ContactService interface
type Server struct {
	db      *sql.DB
	service ContactService
}

// NewServer creates a new server instance
func NewServer(db *sql.DB, service ContactService) *Server {
	return &Server{
		db:      db,
		service: service,
	}
}

// GetContactsHandler handles GET requests to retrieve contacts
func (s *Server) GetContactsHandler(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	contacts, err := s.service.GetContacts(limit, offset)
	if err != nil {
		log.Printf("GetContactsHandler: Error retrieving contacts: %v", err)
		http.Error(w, "Error retrieving contacts", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON200(w, contacts)
}

// AddContactHandler handles POST requests to add a new contact
func (s *Server) AddContactHandler(w http.ResponseWriter, r *http.Request) {
	var contact Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		log.Printf("AddContactHandler: Error decoding request body: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := s.service.AddContact(contact)
	if err != nil {
		log.Printf("AddContactHandler: Error adding contact: %v", err)
		http.Error(w, "Error adding contact", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON200(w, map[string]string{"id": id})
}

// RegisterRoutes registers HTTP routes with the mux router
func (s *Server) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/contacts", s.GetContactsHandler).Methods("GET")
	r.HandleFunc("/contacts", s.AddContactHandler).Methods("POST")
}
