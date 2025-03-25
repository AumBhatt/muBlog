package handlers

import (
	"encoding/json"
	"log"
	"muBlog/internal/models"
	"muBlog/internal/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service}
}

func (handler *UserHandler) GetById(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	user, err := handler.service.GetUserById(ps.ByName("id"))

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

func (handler *UserHandler) CreateUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
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