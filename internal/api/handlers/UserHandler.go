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
	authService *services.AuthService
	userService *services.UserService
}

func NewUserHandler(authService *services.AuthService, userService *services.UserService) *UserHandler {
	return &UserHandler{
		authService: authService,
		userService: userService,
	}
}

func (handler *UserHandler) GetById(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	userId := ps.ByName("id")

	log.Println("/user/", userId)

	if userId == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := handler.userService.GetUserById(userId)

	if err != nil {
		log.Println("Handle GetById:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(res).Encode(schemas.GetUserByIdResponse{
		Id:          user.Id,
		Username:    user.Username,
		MailId:      user.MailId,
		ActiveSince: user.ActiveSince,
	})
}

func (handler *UserHandler) CreateUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	log.Println("/user/new")

	var body schemas.CreateUserRequest
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Println("Handle CreateUser: ", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	validate := validator.New()
	err = validate.Struct(body)

	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(res, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
	}

	response, err := handler.userService.CreateUser(&body)
	if err != nil {
		log.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(&schemas.CreateUserResponse{
			ErrorSchema: &schemas.ErrorSchema{
				Code:    "ErrCreateUser",
				Message: fmt.Sprintf("Error in creating user: %s", err),
			},
		})
		return
	}

	json.NewEncoder(res).Encode(response)
}

func (handler *UserHandler) UserLogin(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	log.Println("/user/login")

	var body schemas.UserLoginRequest
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Println("Handle UserLogin:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	validate := validator.New()
	err = validate.Struct(body)

	if err != nil {
		errors := err.(validator.ValidationErrors)
		json.NewEncoder(res).Encode(schemas.ErrorSchema{
			Code:    "ValidationError",
			Message: errors[len(errors)-1].Field(),
		})
		return
	}

	response, err := handler.authService.CreateToken(&body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(res).Encode(response)
}
