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

func (service *PostService) CreatePost(req schemas.CreatePostRequest) (*schemas.CreatePostResponse, error) {

	post := models.Post{
		Id:        uuid.NewString(),
		CreatedAt: time.Now().UnixMilli(),
		AuthorId:  "",
		Content:   "",
	}

	err := service.postStore.CreatePost(post)

	if err != nil {
		return nil, err
	}

	return &schemas.CreatePostResponse{
		Id: post.Id,
	}, nil
}

func (service *PostService) AddReaction(req schemas.AddReactionRequest) {}
