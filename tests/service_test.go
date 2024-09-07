package tests

import (
	"phonebook-api/contacts"
	"phonebook-api/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test adding a contact to the database
func TestAddContact(t *testing.T) {
	contact := contacts.Contact{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
		Address:   "123 Main St",
	}

	id, err := contacts.AddContact(contact)
	assert.NoError(t, err)

	// Retrieve the contact by ID
	retrievedContact, err := contacts.GetContactByID(id)
	assert.NoError(t, err)

	// Check if the retrieved contact matches the original contact
	assert.Equal(t, contact.FirstName, retrievedContact.FirstName)
	assert.Equal(t, contact.LastName, retrievedContact.LastName)
	assert.Equal(t, contact.Phone, retrievedContact.Phone)
	assert.Equal(t, contact.Address, retrievedContact.Address)
}

// Test retrieving contacts from the database
func TestGetContacts(t *testing.T) {
	contacts, err := contacts.GetContacts(5, 0)
	assert.NoError(t, err)
	assert.NotNil(t, contacts)
	assert.True(t, len(contacts) > 0)
}

// Test updating a contact
func TestUpdateContact(t *testing.T) {
	contact := contacts.Contact{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
		Address:   "123 Main St",
	}

	id, err := contacts.AddContact(contact)
	assert.NoError(t, err)

	updated := contacts.UpdateContactRequest{
		FirstName: utils.Ptr("Jane"),
		LastName:  utils.Ptr("Doe"),
	}

	err = contacts.UpdateContact(id, updated)
	assert.NoError(t, err)

	// Retrieve the contact by ID
	retrievedContact, err := contacts.GetContactByID(id)
	assert.NoError(t, err)

	// Check if the retrieved contact matches the original contact
	assert.Equal(t, *updated.FirstName, retrievedContact.FirstName)
	assert.Equal(t, *updated.LastName, retrievedContact.LastName)
}
