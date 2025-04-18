package stores

import (
	"muBlog/internal/database"
)

type PostStore struct {
	db *database.Connection
}

func NewPostStore(db *database.Connection) *PostStore {
	return &PostStore{db}
}
