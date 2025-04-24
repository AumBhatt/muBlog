package services

import (
	"fmt"
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
		Id:        fmt.Sprintf("post-%s", uuid.NewString()),
		AuthorId:  req.AuthorId,
		Content:   req.Content,
		CreatedAt: time.Now().UnixMilli(),
		EditedAt:  time.Now().UnixMilli(),
	}

	err := service.postStore.CreatePost(post)

	if err != nil {
		return nil, err
	}

	return &schemas.CreatePostResponse{
		Id: post.Id,
	}, nil
}

func (service *PostService) AddReaction(req schemas.AddReactionRequest) (*schemas.AddReactionResponse, error) {

	reaction, err := service.postStore.GetReactionsById(req.PostId)
	if err != nil {
		return nil, err
	}

	if reaction == nil {
		reaction = &models.Reaction{
			Id:        fmt.Sprintf("reaction-%s", uuid.NewString()),
			UserId:    req.UserId,
			PostId:    req.PostId,
			Type:      req.Type,
			CreatedAt: time.Now().UnixMilli(),
			EditedAt:  time.Now().UnixMilli(),
		}

		err = service.postStore.CreateReaction(*reaction)
		if err != nil {
			return nil, err
		}
	} else {
		err = service.postStore.UpdateReaction(reaction.Id, req.Type)
		if err != nil {
			return nil, err
		}
	}

	reactions, err := service.postStore.GetReactionsCountById(req.PostId)
	if err != nil {
		return nil, err
	}

	return &schemas.AddReactionResponse{
		Reactions: reactions,
	}, nil
}

func (service *PostService) GetReactionsCountByPostId(postId string) (*schemas.GetReactionsCountByPostIdResponse, error) {

	reactions, err := service.postStore.GetReactionsCountById(postId)
	if err != nil {
		return nil, err
	}

	return &schemas.GetReactionsCountByPostIdResponse{
		Reactions: reactions,
	}, nil
}

func (service *PostService) GetReactionsPostById(postId string) (*schemas.GetReactionsByPostIdResponse, error) {

	reactions, err := service.postStore.GetReactionsByPostId(postId)
	if err != nil {
		return nil, err
	}

	return &schemas.GetReactionsByPostIdResponse{
		Reactions: reactions,
	}, nil
}
