package database

import (
	"database/sql"
	"log"
	"muBlog/configs"

	_ "github.com/mattn/go-sqlite3"
)

type Connection struct {
	*sql.DB
}

func New() *Connection {
	db, err := sql.Open("sqlite3", configs.DB_PATH)
	if err != nil {
		log.Fatalln("Database: new connection failedn\n", err)
	}

	return &Connection{db}
}
