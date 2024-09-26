package contacts

import (
	"log"
)

// GetContacts retrieves contacts from the database
func (s *Service) GetContacts(limit, offset int) ([]Contact, error) {
	log.Printf("Service: Retrieving contacts with limit %d, offset %d", limit, offset)

	rows, err := s.db.Query("SELECT id, first_name, last_name, phone, address FROM contacts LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []Contact
	for rows.Next() {
		var c Contact
		if err := rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Phone, &c.Address); err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}
	return contacts, nil
}

// AddContact adds a new contact to the database
func (s *Service) AddContact(contact Contact) (string, error) {
	log.Printf("Service: Adding new contact: %+v", contact)

	var id string
	err := s.db.QueryRow(
		"INSERT INTO contacts (first_name, last_name, phone, address) VALUES ($1, $2, $3, $4) RETURNING id",
		contact.FirstName, contact.LastName, contact.Phone, contact.Address).Scan(&id)

	if err != nil {
		return "", err
	}
	return id, nil
}
