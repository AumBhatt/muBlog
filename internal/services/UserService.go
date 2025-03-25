package services

import (
	"muBlog/internal/api/schemas"
	"muBlog/internal/models"
	"muBlog/internal/stores"
	"time"

	"github.com/google/uuid"
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

func (service *UserService) CreateUser(request *schemas.CreateUserRequest) (*models.User, error) {

	user := &models.User{
		Id: uuid.NewString(),
		Username: request.Username,
		MailId: request.MailId,
		ActiveSince: time.Now().UnixMilli(),
	}

	err := service.store.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
