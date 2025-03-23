package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Connection struct {
	*sql.DB
}

func New() *Connection {
	db, err := sql.Open("sqlite3", "/home/aum/learn/go/muBlog/internal/database/mu")
	if err != nil {
		log.Fatalln("Database: new connection failedn\n", err)
	}

	return &Connection{db}
}
