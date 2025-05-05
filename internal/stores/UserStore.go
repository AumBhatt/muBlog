package stores

import (
	"fmt"
	"muBlog/internal/database"
	"muBlog/internal/models"
)

type UserStore struct {
	db *database.Connection
}

func NewUserStore(db *database.Connection) *UserStore {
	return &UserStore{db}
}

func (store *UserStore) FindById(id string) (*models.User, error) {

	stmt, err := store.db.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("FindById UserStore: %s", err)
	}

	user := &models.User{}
	row := stmt.QueryRow(id)
	row.Scan(&user.Id, &user.Username, &user.Email, &user.ActiveSince, &user.Password)

	return user, nil
}

func (store *UserStore) FindByUsername(username string) (*models.User, error) {

	stmt, err := store.db.Prepare("SELECT * FROM users WHERE username = ?")
	if err != nil {
		return nil, fmt.Errorf("FindBy UserStore: %s", err)
	}

	user := &models.User{}
	row := stmt.QueryRow(username)
	row.Scan(&user.Id, &user.Username, &user.Email, &user.ActiveSince, &user.Password)

	return user, nil
}

func (store *UserStore) CreateUser(user *models.User) error {

	stmt, err := store.db.Prepare("INSERT INTO users (id, username, email, active_since, password) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("AddUser UserStore: %s", err)
	}

	_, err = stmt.Exec(user.Id, user.Username, user.Email, user.ActiveSince, user.Password)
	if err != nil {
		return fmt.Errorf("AddUser UserStore: %s", err)
	}

	return nil
}

func (store *UserStore) AddFollower(follow *models.Follow) error {

	stmt, err := store.db.Prepare("INSERT INTO follow (id, user_id, followers_id) VALUES (?, ?, ?)")
	if err != nil {
		return fmt.Errorf("UserStore.AddFollower: %s", err)
	}

	_, err = stmt.Exec(follow.Id, follow.UserId, follow.FollowerId)
	if err != nil {
		return fmt.Errorf("UserStore.AddFollower: %s", err)
	}

	return nil
}

func (store *UserStore) RemoveFollower(userId string, followerId string) error {

	stmt, err := store.db.Prepare(`
		DELETE FROM follow
			WHERE userId = ? AND followerId = ?
	`)
	if err != nil {
		return fmt.Errorf("UserStore.AddFollower: %s", err)
	}

	_, err = stmt.Exec(userId, followerId)
	if err != nil {
		return fmt.Errorf("UserStore.RemoveFollower: %s", err)
	}

	return nil
}
