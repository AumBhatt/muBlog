package services

import (
	"muBlog/internal/api/schemas"
	"muBlog/internal/models"
	"muBlog/internal/stores"
	"time"

	"github.com/google/uuid"
)

type PostService struct {
	postStore *stores.PostStore
	userStore *stores.UserStore
}

func NewPostService(postStore *stores.PostStore, userStore *stores.UserStore) *PostService {
	return &PostService{postStore, userStore}
}

func (service *PostService) CreatePost(res schemas.CreatePostResponse, req schemas.CreatePostRequest) error {

	service.postStore.CreatePost(models.Post{
		Id:        uuid.NewString(),
		CreatedAt: time.Now().UnixMilli(),
		AuthorId:  "",
		Content:   "",
		Reactions: map[models.Reaction][]string{},
	})
	return nil
}
