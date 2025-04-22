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
		Id:         uuid.NewString(),
		CreatedAt:  time.Now().UnixMilli(),
		AuthorId:   req.AuthorId,
		Content:    req.Content,
		ReactionId: nil,
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
			Id:        uuid.NewString(),
			UserId:    req.UserId,
			Type:      req.ReactionType,
			Timestamp: time.Now().UnixMilli(),
		}

		err = service.postStore.CreateReaction(*reaction)

		if err != nil {
			return nil, err
		}

		response, err := service.postStore.GetUsersByReactions(reaction.Id)
		if err != nil {
			return nil, err
		}

		return &schemas.AddReactionResponse{
			Reactions: response,
		}, nil
	}

	err = service.postStore.UpdateReaction(reaction.Id, req.ReactionType)
	if err != nil {
		return nil, err
	}

	response, err := service.postStore.GetUsersByReactions(reaction.Id)
	if err != nil {
		return nil, err
	}

	return &schemas.AddReactionResponse{
		Reactions: response,
	}, nil
}
