package users

import (
	"fmt"

	"github.com/lib/pq"
)

// CreateUser inserts a new user into the database and returns the user details
func CreateUser(params CreateUserParams) (*User, error) {
	// Define the SQL query for inserting a user and returning the columns
	query := `
    INSERT INTO users (
        username,
        hashed_password,
        full_name,
        email
    ) VALUES (
        $1, $2, $3, $4
    ) RETURNING username, hashed_password, full_name, email, created_at;`

	// Execute the SQL query and retrieve the result
	var user User
	err := DB.QueryRow(query, params.Username, params.HashedPassword, params.FullName, params.Email).
		Scan(&user.Username, &user.HashedPassword, &user.FullName, &user.Email, &user.CreatedAt)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, fmt.Errorf("unique constraint violation: %w", err)
			}
		}
		return nil, fmt.Errorf("error inserting user: %w", err)
	}

	return &user, nil
}

func GetUser(username string) (User, error) {
	const query = `SELECT username, hashed_password, full_name, email, created_at FROM users WHERE username = $1`
	row := DB.QueryRow(query, username)
	var user User
	err := row.Scan(
		&user.Username,
		&user.HashedPassword,
		&user.FullName,
		&user.Email,
		&user.CreatedAt,
	)
	return user, err
}
