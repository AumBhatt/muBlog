package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"muBlog/internal/api/schemas"
	"muBlog/internal/services"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service}
}

func (handler *UserHandler) GetById(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	log.Println("/users/", ps.ByName("id"))

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

	log.Println("/users/new")

	var body schemas.CreateUserRequest
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Println("HandleAddUser: ", err)
		res.WriteHeader(500)
		return
	}

	validate := validator.New()
	err = validate.Struct(body)

	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(res, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
	}

	user, err := handler.service.CreateUser(&body)
	if err != nil {
		log.Println(err)
		res.WriteHeader(500)
		json.NewEncoder(res).Encode(&schemas.CreateUserResponse{
			ErrorSchema: &schemas.ErrorSchema{
				Code: "ErrCreateUser",
				Message: fmt.Sprintf("Error in creating user: %s", err),
			},
		})
		return
	}

	json.NewEncoder(res).Encode(&schemas.CreateUserResponse{
		Id: &user.Id,
		Username: &user.Username,
	})
}