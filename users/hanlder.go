package users

import (
	"encoding/json"
	"log"
	"net/http"
	"phonebook-api/server"
	"phonebook-api/utils"
	"time"

	"github.com/gorilla/mux"
)

// var DB *sql.DB
// var TM token.Maker

// // SetDB sets the database connection
// func SetDB(db *sql.DB) {
// 	DB = db
// }

// func SetTokenMaker(tokenMaker token.Maker) {
// 	TM = tokenMaker
// }

// RegisterRoutes sets up the HTTP routes for contacts
func RegisterRoutes(r *mux.Router, server *server.Server) {

	r.HandleFunc("/users", server.createUserHandler).Methods("POST")
	r.HandleFunc("/users/login", server.LoginUserHandler).Methods("POST")
}

// createUser handles the creation of a new user
func (s *server.Server) createUserHandler(w http.ResponseWriter, r *http.Request) {
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

	user, err := s.CreateUser(cup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON200(w, newUserResponse(*user))
}

func newUserResponse(user User) UserResponse {
	return UserResponse{
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

// loginUser handles user login and token generation
func (s *server.Server) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
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

	retrievedUser, err := s.GetUser(reqUser.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = utils.CheckPassword(reqUser.Password, retrievedUser.HashedPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	accessToken, _, err := s.CreateToken(
		retrievedUser.Username,
		time.Minute*15,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := &LoginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(retrievedUser),
	}

	utils.WriteJSON200(w, res)
}
