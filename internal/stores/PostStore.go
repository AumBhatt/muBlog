package stores

import (
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
		return fmt.Errorf("CreatePost db err: %s", err)
	}

	_, err = stmt.Exec(post.Id, post.CreatedAt, post.AuthorId, post.Content, post.ReactionId)
	if err != nil {
		return fmt.Errorf("CreatePost err: %s", err)
	}

	return nil
}

func (store *PostStore) GetPostReactionsById(id string) (*models.Reactions, error) {

	stmt, err := store.db.Prepare("SELECT * FROM reactions WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("GetPostreactionsById err: %s", err)
	}

	var rowData *models.Reactions

	stmt.QueryRow(id).Scan(rowData)

	return rowData, nil
}

func (store *PostStore) CreateReaction(reaction models.Reactions) error {

	stmt, err := store.db.Prepare("INSERT INTO reactions (ids, userId, type, timestamp) VALUES (?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("AddReaction err: %s", err)
	}

	_, err = stmt.Exec(reaction.Id, reaction.UserId, reaction.Type, reaction.Timestamp)
	if err != nil {
		return fmt.Errorf("AddReaction err: %s", err)
	}

	return nil
}

func (store *PostStore) UpdateReaction(reaction models.Reactions) error {
	return nil
}
