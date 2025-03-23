package stores

import (
	"log"
	"muBlog/internal/database"
	"muBlog/internal/models"
)

type UserStore struct {
	db *database.Connection
}

func NewUserStore(db *database.Connection) *UserStore {
	return &UserStore{db}
}

func (store *UserStore) FindById(id string) *models.User {
	stmt, err := store.db.Prepare("SELECT * FROM users WHERE id = ?")

	if err != nil {
		log.Fatalln(err)
	}

	user := &models.User{}
	row := stmt.QueryRow(id)
	row.Scan(&user.Id, &user.Username, &user.MailId, &user.ActiveSince)

	return user
}
