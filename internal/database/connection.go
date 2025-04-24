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

type DatabaseInitFunc func(*Connection) error

func New(initFunc DatabaseInitFunc) *Connection {
	db, err := sql.Open("sqlite3", configs.DB_PATH)
	if err != nil {
		log.Fatalln("Database connection err:", err)
	}

	conn := &Connection{db}

	err = initFunc(conn)
	if err != nil {
		log.Fatalln(err)
	}

	return conn
}
