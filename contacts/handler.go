package contacts

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var DB *sql.DB

// SetDB sets the database connection
func SetDB(db *sql.DB) {
	DB = db
}

// RegisterRoutes sets up the HTTP routes for contacts
func RegisterRoutes(r *mux.Router, db *sql.DB) {
	SetDB(db)
	r.HandleFunc("/contacts", GetContactsHandler).Methods("GET")
	r.HandleFunc("/contacts", AddContactHandler).Methods("POST")
	r.HandleFunc("/contacts", UpdateContactHandler).Methods("PUT")
	r.HandleFunc("/contacts/{id}", DeleteContactHandler).Methods("DELETE")
	r.HandleFunc("/contacts/search", SearchContactHandler).Methods("GET")
}

const (
	MaxLimit = 10 // maximum number of items per page
)

// GetContactsHandler handles GET requests for contacts with pagination
func GetContactsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetContactsHandler: Handling GET /contacts request")

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 || limit > MaxLimit {
		log.Printf("GetContactsHandler: Invalid or missing limit, using default limit of %d", MaxLimit)
		limit = MaxLimit // enforce default limit
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		log.Println("GetContactsHandler: Invalid or missing offset, using default offset of 0")
		offset = 0 // default offset
	}

	contacts, err := GetContacts(limit, offset)
	if err != nil {
		log.Printf("GetContactsHandler: Error fetching contacts: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("GetContactsHandler: Fetched %d contacts", len(contacts))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

// SearchContactHandler handles GET requests to search contacts
func SearchContactHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("SearchContactHandler: Handling GET /contacts/search request")

	term := r.URL.Query().Get("term")
	if term == "" {
		log.Println("SearchContactHandler: Search term is missing")
		http.Error(w, "Search term is required", http.StatusBadRequest)
		return
	}

	contacts, err := SearchContacts(term)
	if err != nil {
		log.Printf("SearchContactHandler: Error searching contacts: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("SearchContactHandler: Found %d contacts for term: %s", len(contacts), term)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

// AddContactHandler handles POST requests to add a new contact
func AddContactHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AddContactHandler: Handling POST /contacts request")
	var contact Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		log.Printf("AddContactHandler: Error decoding contact: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := AddContact(contact)
	if err != nil {
		log.Printf("AddContactHandler: Error adding contact: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("AddContactHandler: Added new contact with ID: %s", id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contact)
}

// UpdateContactHandler handles PUT requests to update an existing contact
func UpdateContactHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateContactHandler: Handling PUT /contacts request")
	var updatedContact UpdateContactRequest
	if err := json.NewDecoder(r.Body).Decode(&updatedContact); err != nil {
		log.Printf("UpdateContactHandler: Error decoding contact update: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if updatedContact.ID == nil {
		log.Println("UpdateContactHandler: Missing contact ID for update")
		http.Error(w, "Update handler must get contact ID in order to update the relevant item.", http.StatusNotFound)
		return
	}

	id := *updatedContact.ID
	err := UpdateContact(id, updatedContact)
	if err != nil {
		log.Printf("UpdateContactHandler: Error updating contact with ID %s: %v", id, err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	log.Printf("UpdateContactHandler: Updated contact with ID: %s", id)
	w.WriteHeader(http.StatusOK)
}

// DeleteContactHandler handles DELETE requests to remove a contact
func DeleteContactHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteContactHandler: Handling DELETE /contacts/{id} request")
	params := mux.Vars(r)
	id := params["id"]
	err := DeleteContact(id)
	if err != nil {
		log.Printf("DeleteContactHandler: Error deleting contact with ID %s: %v", id, err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	log.Printf("DeleteContactHandler: Deleted contact with ID: %s", id)
	w.WriteHeader(http.StatusOK)
}
