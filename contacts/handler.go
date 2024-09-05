package contacts

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var DB *sql.DB

// RegisterRoutes sets up the HTTP routes for contacts
func RegisterRoutes(r *mux.Router, db *sql.DB) {
	DB = db
	r.HandleFunc("/contacts", GetContactsHandler).Methods("GET")
	r.HandleFunc("/contacts", AddContactHandler).Methods("POST")
	r.HandleFunc("/contacts/{id}", UpdateContactHandler).Methods("PUT")
	r.HandleFunc("/contacts/{id}", DeleteContactHandler).Methods("DELETE")
	r.HandleFunc("/contacts/search", SearchContactHandler).Methods("GET")
}

// GetContactsHandler handles GET requests for contacts with pagination
func GetContactsHandler(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10 // default limit
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0 // default offset
	}

	contacts, err := GetContacts(limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

// SearchContactHandler handles GET requests to search contacts
func SearchContactHandler(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("term")
	if term == "" {
		http.Error(w, "Search term is required", http.StatusBadRequest)
		return
	}

	contacts, err := SearchContacts(term)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

// AddContactHandler handles POST requests to add a new contact
func AddContactHandler(w http.ResponseWriter, r *http.Request) {
	var contact Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := AddContact(contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contact)
}

// UpdateContactHandler handles PUT requests to update an existing contact
func UpdateContactHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedContact UpdateContactRequest
	if err := json.NewDecoder(r.Body).Decode(&updatedContact); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := UpdateContact(params["id"], updatedContact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteContactHandler handles DELETE requests to remove a contact
func DeleteContactHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := DeleteContact(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
