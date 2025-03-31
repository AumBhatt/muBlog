package services

import (
	"fmt"
	"log"
	"muBlog/internal/api/schemas"
	"muBlog/internal/services/utils"
	"muBlog/internal/stores"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	secret    []byte
	userStore *stores.UserStore
}

func NewAuthService(userStore *stores.UserStore) *AuthService {
	secret, err := utils.GenerateNewSecret()
	if err != nil {
		log.Fatalln(err)
	}

	return &AuthService{
		secret:    secret,
		userStore: userStore,
	}
}

func (service *AuthService) CreateToken(request *schemas.UserLoginRequest) (*schemas.UserLoginResponse, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if err != nil {
		return nil, err
	}

	user, err := service.userStore.FindByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return &schemas.UserLoginResponse{
			ErrorSchema: &schemas.ErrorSchema{
				Code:    "NoUserWithUsername",
				Message: fmt.Sprintf("No user with username '%s' found", user.Username),
			},
		}, nil
	}

	if string(hashedPassword) != user.Password {
		return &schemas.UserLoginResponse{
			ErrorSchema: &schemas.ErrorSchema{
				Code:    "InvalidPassword",
				Message: "Authentication failed due to invalid password.",
			},
		}, nil
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		jwt.MapClaims{
			"username": request.Username,
		},
	)

	tokenString, err := token.SignedString(service.secret)
	if err != nil {
		return nil, err
	}

	return &schemas.UserLoginResponse{
		Id:    &user.Id,
		Token: &tokenString,
	}, nil
}

func (service *AuthService) VerifyToken(tokenString string) (*schemas.ErrorSchema, error) {
	token, err := jwt.Parse(tokenString[len("Bearer "):], service.authKeyFunc)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return &schemas.ErrorSchema{
			Code:    "InvalidToken",
			Message: "Authentication failed due to invalid token.",
		}, nil
	}

	return nil, nil
}

func (service *AuthService) authKeyFunc(token *jwt.Token) (interface{}, error) {
	return service.secret, nil
}
