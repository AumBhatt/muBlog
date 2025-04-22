package stores

import (
	"database/sql"
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

	stmt, err := store.db.Prepare(`
		INSERT INTO posts (id, createdAt, authorId, content, reactionId) 
					VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("CreatePost db err: %s", err)
	}

	_, err = stmt.Exec(post.Id, post.CreatedAt, post.AuthorId, post.Content, post.ReactionId)
	if err != nil {
		return fmt.Errorf("CreatePost err: %s", err)
	}

	return nil
}

func (store *PostStore) GetReactionsById(id string) (*models.Reaction, error) {

	stmt, err := store.db.Prepare("SELECT * FROM reactions WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("GetPostreactionsById err: %s", err)
	}

	var rowData *models.Reaction

	err = stmt.QueryRow(id).Scan(rowData)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("GetPostreactionsById err: %s", err)
	}

	return rowData, nil
}

func (store *PostStore) CreateReaction(reaction models.Reaction) error {

	stmt, err := store.db.Prepare(`
		INSERT INTO reactions (id, userId, type, timestamp) 
					VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("CreateReaction err: %s", err)
	}

	_, err = stmt.Exec(reaction.Id, reaction.UserId, reaction.Type, reaction.Timestamp)
	if err != nil {
		return fmt.Errorf("CreateReaction err: %s", err)
	}

	return nil
}

func (store *PostStore) UpdateReaction(reactionId string, reactionType string) error {

	stmt, err := store.db.Prepare(`
		UPDATE reactions
			SET type = ?
			WHERE id = ?
	`)
	if err != nil {
		return fmt.Errorf("UpdateReaction err: %s", err)
	}

	_, err = stmt.Exec(reactionType, reactionId)
	if err != nil {
		return fmt.Errorf("UpdateReaction err: %s", err)
	}
	return nil
}

func (store *PostStore) GetUsersByReactions(reactionId string) ([]map[string]string, error) {

	var data []map[string]string
	stmt, err := store.db.Prepare(`
		SELECT users.id as "userId", users.username as "username", reactions.type as "type"
			FROM users
			INNER JOIN reactions ON
			users.id = reactions.userId
			WHERE reactions.Id = ?
	`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(reactionId)
	if err != nil {
		return nil, fmt.Errorf("UpdateReaction err: %s", err)
	}

	for rows.Next() {
		var userId string
		var username string
		var reactionType string

		err = rows.Scan(&userId, &username, &reactionType)
		if err != nil {
			return nil, err
		}

		data = append(data, map[string]string{
			"userId":   userId,
			"username": username,
			"type":     reactionType,
		})
	}

	return data, nil
}
