package services

import (
	"muBlog/internal/api/schemas"
	"muBlog/internal/models"
	"muBlog/internal/stores"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

type UserService struct {
	store *stores.UserStore
}

func NewUserService(store *stores.UserStore) *UserService {
	return &UserService{store}
}

func (service *UserService) GetUserById(id string) (*schemas.GetUserByIdResponse, error) {
	user, err := service.store.FindById(id)
	if err != nil {
		return nil, err
	}

	return &schemas.GetUserByIdResponse{
		Id:          user.Id,
		Username:    user.Username,
		Email:       user.Email,
		ActiveSince: user.ActiveSince,
	}, nil
}

func (service *UserService) CreateUser(request *schemas.SignupRequest) (*schemas.SignupResponse, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Id:          uuid.NewString(),
		Username:    request.Username,
		Email:       request.Email,
		ActiveSince: time.Now().UnixMilli(),
		Password:    string(hashedPassword),
	}

	err = service.store.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &schemas.SignupResponse{
		Id:       &user.Id,
		Username: &user.Username,
	}, nil
}
