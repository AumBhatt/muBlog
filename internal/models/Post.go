package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id        string
	CreatedAt int64
	Content   string
	Likes     []string
	UserId    string
}

func NewPost(userId string, content string) *Post {
	return &Post{
		Id:        uuid.New().String(),
		CreatedAt: time.Time.UnixMilli(time.Now()),
		Content:   content,
		Likes:     []string{},
		UserId:    userId,
	}
}
