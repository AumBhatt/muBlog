package handlers

import (
	"encoding/json"
	"log"
	"muBlog/internal/models"
	"muBlog/internal/services"
	"net/http"
	"strings"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service}
}

func (handler *UserHandler) GetById(res http.ResponseWriter, req *http.Request) {
	paths := strings.Split(req.URL.Path, "/users/")
	id := paths[len(paths)-1]

	user, err := handler.service.GetUserById(id)

	if err != nil {
		log.Println(err)
		res.WriteHeader(500)
		return
	}

	payload, err := json.Marshal(user)
	if err != nil {
		log.Println("HandleGetById: ", err)
		res.WriteHeader(500)
		return
	}

	_, err = res.Write(payload)
	if err != nil {
		log.Println("HandleGetById: ", err)
		res.WriteHeader(500)
		return
	}
}

func (handler *UserHandler) CreateUser(res http.ResponseWriter, req *http.Request) {
	var newUser *models.User
	err := json.NewDecoder(req.Body).Decode(newUser)
	if err != nil {
		log.Println("HandleAddUser: ", err)
		res.WriteHeader(500)
		return
	}

	err = handler.service.CreateUser(newUser)
	if err != nil {
		log.Println(err)
		res.WriteHeader(500)
		return
	}
}