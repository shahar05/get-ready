package server

import (
	"database/sql"
	"phonebook-api/token"
)

type Server struct {
	db         *sql.DB
	tokenMaker token.Maker
	// TODO: add here some configs ( config Config )
}
