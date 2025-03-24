package services

import (
	"muBlog/internal/models"
	"muBlog/internal/stores"
)

type UserService struct {
	store *stores.UserStore
}

func NewUserService(store *stores.UserStore) *UserService {
	return &UserService{store}
}

func (service *UserService) GetUserById(id string) (*models.User, error) {
	return service.store.FindById(id)
}

func (service *UserService) CreateUser(user *models.User) error {
	return service.store.CreateUser(user)
}
