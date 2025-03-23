package handlers

import (
	"encoding/json"
	"log"
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

func (handler *UserHandler) HandleGetById(res http.ResponseWriter, req *http.Request) {
	paths := strings.Split(req.URL.Path, "/users/")
	id := paths[len(paths)-1]

	user := handler.service.GetUserById(id)
	payload, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	res.Write(payload)
}
