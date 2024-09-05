package contacts

// Contact model
type Contact struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

// Contact represents a contact with optional fields
type UpdateContactRequest struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Phone     *string `json:"phone"`
	Address   *string `json:"address"`
}
