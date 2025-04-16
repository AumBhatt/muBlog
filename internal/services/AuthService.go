package services

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
	"muBlog/internal/api/schemas"
	"muBlog/internal/stores"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	secret    *ecdsa.PrivateKey
	userStore *stores.UserStore
}

func NewAuthService(userStore *stores.UserStore) *AuthService {
	secret, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalln(err)
	}

	return &AuthService{
		secret:    secret,
		userStore: userStore,
	}
}

func (service *AuthService) CreateToken(request *schemas.LoginRequest) (*schemas.LoginResponse, error) {

	user, err := service.userStore.FindByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return &schemas.LoginResponse{
			ErrorSchema: &schemas.ErrorSchema{
				Code:    "NoUserWithUsername",
				Message: fmt.Sprintf("No user with username '%s' found", request.Username),
			},
		}, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		return &schemas.LoginResponse{
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
		log.Fatalln("Error generating token", err)
		return nil, err
	}

	return &schemas.LoginResponse{
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
