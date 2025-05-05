package stores

import (
	"fmt"
	"muBlog/internal/database"
)

func InitStores(conn *database.Connection) error {

	_, err := conn.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY NOT NULL,
			username TEXT NOT NULL,
			email TEXT NOT NULL,
			active_since INTEGER NOT NULL,
			password TEXT NOT NULL
		)
	`)
	if err != nil {
		return fmt.Errorf("InitStore - users table creation err: %s", err)
	}

	_, err = conn.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id TEXT PRIMARY KEY NOT NULL,
			author_id TEXT NOT NULL,
			content TEXT,
			created_at INTEGER NOT NULL,
			edited_at INTEGER,
			FOREIGN KEY (author_id) REFERENCES users(id)

		)
	`)
	if err != nil {
		return fmt.Errorf("InitStore - posts table creation err: %s", err)
	}

	_, err = conn.Exec(`
		CREATE TABLE IF NOT EXISTS reactions (
			id TEXT PRIMARY KEY NOT NULL,
			user_id TEXT NOT NULL,
			post_id TEXT NOT NULL,
			type TEXT NOT NULL
				CHECK (type IN ('like', 'dislike', 'funny', 'support')),
			created_at INTEGER NOT NULL,
			edited_at INTEGER,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (post_id) REFERENCES posts(id)
		)
	`)
	if err != nil {
		return fmt.Errorf("InitStore - reactions table creation err: %s", err)
	}

	_, err = conn.Exec(`
		CREATE TABLE IF NOT EXISTS follow (
			id TEXT PRIMARY KEY NOT NULL,
			user_id TEXT NOT NULL,
			followers_id TEXT NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (followers_id) REFERENCES users(id)
		)
	`)
	if err != nil {
		return fmt.Errorf("InitStore - reactions table creation err: %s", err)
	}

	return nil
}
