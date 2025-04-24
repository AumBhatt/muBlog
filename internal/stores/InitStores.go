package stores

import (
	"fmt"
	"muBlog/internal/database"
)

func InitStores(conn *database.Connection) error {

	_, err := conn.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY NOT NULL
				CHECK (id LIKE 'user-%'),
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
			id TEXT PRIMARY KEY NOT NULL
				CHECK (id LIKE 'post-%'),
			author_id TEXT NOT NULL
				CHECK (author_id LIKE 'user-%'),
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
			id TEXT PRIMARY KEY NOT NULL
				CHECK (id LIKE 'reaction-%'),
			user_id TEXT NOT NULL
				CHECK (user_id LIKE 'user-%'),
			post_id TEXT NOT NULL
				CHECK (post_id LIKE 'post-%'),
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

	return nil
}
