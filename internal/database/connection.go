package database

import (
	"database/sql"
	"log"
)

type Connection struct {
	*sql.DB
}

func New() *Connection {
	db, err := sql.Open("sqlite3", "./mu.db")
	if err != nil {
		log.Fatalln("Database: new connection failedn\n", err)
	}

	return &Connection{db}
}
