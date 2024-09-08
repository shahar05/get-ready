package users

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"phonebook-api/utils"

	"github.com/gorilla/mux"
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/util"
)

var DB *sql.DB

// SetDB sets the database connection
func SetDB(db *sql.DB) {
	DB = db
}

// RegisterRoutes sets up the HTTP routes for contacts
func RegisterRoutes(r *mux.Router, db *sql.DB) {
	SetDB(db)
	r.HandleFunc("/users", createUserHandler).Methods("POST")
	r.HandleFunc("/users/login", LoginUserHandler).Methods("POST")
}

// createUser handles the creation of a new user
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cup := CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := CreateUser(cup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON200(w, newUserResponse(user))
}

func newUserResponse(user db.User) UserResponse {
	return UserResponse{
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

// loginUser handles user login and token generation
func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	type loginUserRequest struct {
		Username string
		Password string
	}

	var reqUser loginUserRequest

	if err := json.NewDecoder(r.Body).Decode(&reqUser); err != nil {
		log.Printf("AddContactHandler: Error decoding contact: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	retrievedUser, err := GetUser(reqUser.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = util.CheckPassword(reqUser.Password, retrievedUser.HashedPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	accessToken, _, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// res := &loginUserResponse{
	// 	AccessToken: accessToken,
	// 	User:        newUserResponse(user),
	// }

	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(res)
}
