package stores

import (
	"encoding/json"
	"fmt"
	"muBlog/internal/database"
	"muBlog/internal/models"
)

type PostStore struct {
	db *database.Connection
}

func NewPostStore(db *database.Connection) *PostStore {
	return &PostStore{db}
}

func (store *PostStore) CreatePost(post models.Post) error {

	stmt, err := store.db.Prepare("INSERT INTO posts (ids, createdAt, authorId, content, reactions) VALUES (?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("CreatePost db err:", err)
	}

	reactionsMarshalled, _ := json.Marshal(post.Reactions)
	if err != nil {
		return fmt.Errorf("CreatePost json marshalling err", err)
	}

	_, err = stmt.Exec(post.Id, post.CreatedAt, post.AuthorId, post.Content, string(reactionsMarshalled))
	if err != nil {
		return fmt.Errorf("CreatePost err:", err)
	}

	return nil
}
