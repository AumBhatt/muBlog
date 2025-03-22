package stores

import (
	"log"
	"muBlog/internal/database"
	"muBlog/internal/models"
)

type UserStore struct {
	db *database.Connection
}

func (store *UserStore) FindById(id string) *models.User {
	stmt, err := store.db.Prepare("SELECT * FROM users WHERE id = ?")

	if err != nil {
		log.Fatalln(err)
	}

	user := &models.User{}
	row := stmt.QueryRow(stmt)
	row.Scan(&user)

	return user
}
