package services

import (
	"muBlog/internal/models"
	"muBlog/internal/stores"
)

type UserService struct {
	store *stores.UserStore
}

func (service *UserService) GetUserById(id string) *models.User {
	return service.store.FindById(id)
}
