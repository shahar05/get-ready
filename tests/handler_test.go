package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"phonebook-api/contacts"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test GET contacts with pagination
func TestGetContactsHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/contacts?limit=5&offset=0", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(contacts.GetContactsHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

// Test POST new contact
func TestAddContactHandler(t *testing.T) {
	contactJSON := `{"first_name":"John","last_name":"Doe","phone":"1234567890","address":"123 Main St"}`
	req, _ := http.NewRequest("POST", "/contacts", strings.NewReader(contactJSON))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(contacts.AddContactHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

// Test PUT update contact
func TestUpdateContactHandler(t *testing.T) {
	// Step 1: Add a contact to get an ID
	contact := contacts.Contact{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
		Address:   "123 Main St",
	}
	id, err := contacts.AddContact(contact)
	if err != nil {
		t.Fatalf("Failed to add contact: %v", err)
	}

	// Step 2: updated the contact
	expectedContact := contacts.Contact{
		FirstName: "Jane",
		LastName:  "DoeDoeDoe",
		ID:        id,
	}

	expectedContactBytes, err := json.Marshal(expectedContact)
	if err != nil {
		t.Fatalf("Failed to Marshal expectedContact: %v", err)
	}

	req, err := http.NewRequest("PUT", "/contacts", strings.NewReader(string(expectedContactBytes)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(contacts.UpdateContactHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Step 3: Retrieve the contact from the database
	actualContact, err := contacts.GetContactByID(id)
	if err != nil {
		t.Fatalf("Failed to retrieve contact: %v", err)
	}

	// Step 4: Assert that the contact has been updated correctly
	assert.Equal(t, expectedContact.FirstName, actualContact.FirstName)
	assert.Equal(t, expectedContact.LastName, actualContact.LastName)
}
