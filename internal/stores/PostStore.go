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
		INSERT INTO posts (
			id, 
			author_id, 
			content, 
			created_at, 
			edited_at
		) 
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("PostStore.CreatePost prepare err: %s", err)
	}

	_, err = stmt.Exec(
		post.Id,
		post.AuthorId,
		post.Content,
		post.CreatedAt,
		post.EditedAt,
	)
	if err != nil {
		return fmt.Errorf("PostStore.CreatePost exec err: %s", err)
	}

	return nil
}

func (store *PostStore) GetPostById(postId string) (*models.Post, error) {

	stmt, err := store.db.Prepare(`
		SELECT * FROM posts
			WHERE id = ?
	`)
	if err != nil {
		return nil, fmt.Errorf("PostStore.GetPostById prepare err: %s", err)
	}

	row := stmt.QueryRow(postId)
	var post models.Post
	err = row.Scan(&post.Id, &post.AuthorId, &post.Content, &post.CreatedAt, &post.EditedAt)
	if err != nil {
		return nil, fmt.Errorf("PostStore.GetPostById prepare err: %s", err)
	}

	return &post, nil
}

func (store *PostStore) GetReactionsById(id string) (*models.Reaction, error) {

	stmt, err := store.db.Prepare("SELECT * FROM reactions WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("PostStore.GetPostreactionsById err: %s", err)
	}

	var rowData *models.Reaction

	err = stmt.QueryRow(id).Scan(rowData)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("PostStore.GetReactionsById err: %s", err)
	}

	return rowData, nil
}

func (store *PostStore) CreateReaction(reaction models.Reaction) error {

	stmt, err := store.db.Prepare(`
		INSERT INTO reactions (
			id, 
			user_id, 
			post_id, 
			type, 
			created_at, 
			edited_at
		) 
		VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("PostStore.CreateReaction err: %s", err)
	}

	_, err = stmt.Exec(
		reaction.Id,
		reaction.UserId,
		reaction.PostId,
		reaction.Type,
		reaction.CreatedAt,
		reaction.EditedAt)
	if err != nil {
		return fmt.Errorf("PostStore.CreateReaction err: %s", err)
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
		return fmt.Errorf("PostStore.UpdateReaction err: %s", err)
	}

	_, err = stmt.Exec(reactionType, reactionId)
	if err != nil {
		return fmt.Errorf("PostStore.UpdateReaction err: %s", err)
	}
	return nil
}

func (store *PostStore) GetReactionsByPostId(postId string) ([]map[string]string, error) {

	var data []map[string]string
	stmt, err := store.db.Prepare(`
		SELECT users.id as "userId", users.username as "username", reactions.type as "type"
			FROM users
			INNER JOIN reactions ON
			users.id = reactions.user_id
			WHERE post_id = ?
	`)
	if err != nil {
		return nil, fmt.Errorf("PostStore.GetUsersByReactions err: %s", err)
	}

	rows, err := stmt.Query(postId)
	if err != nil {
		return nil, fmt.Errorf("PostStore.GetUsersByReactions err: %s", err)
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

func (store *PostStore) GetReactionsCountById(postId string) ([]map[string]any, error) {

	var data []map[string]any

	stmt, err := store.db.Prepare(`
		SELECT type, COUNT(*)
			FROM reactions
			WHERE post_id = ?
			GROUP BY type
	`)
	if err != nil {
		return nil, fmt.Errorf("PostStore.GetReactionsCountById err: %s", err)
	}

	rows, err := stmt.Query(postId)
	if err != nil {
		return nil, fmt.Errorf("PostStore.GetReactionsCountById err: %s", err)
	}

	for rows.Next() {
		var reactionType string
		var count int

		err = rows.Scan(&reactionType, &count)
		if err != nil {
			return nil, err
		}

		row := map[string]any{
			"type":  reactionType,
			"count": count,
		}
		data = append(data, row)
	}

	return data, nil
}
