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

type AuthHandler struct {
	authService *services.AuthService
	userService *services.UserService
}

func NewAuthHandler(authService *services.AuthService, userService *services.UserService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userService: userService,
	}
}

func (handler *AuthHandler) Signup(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	log.Println("/auth/signup")

	var body schemas.SignupRequest
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Println("Handle CreateUser: ", err.Error())
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
		json.NewEncoder(res).Encode(&schemas.SignupResponse{
			ErrorSchema: &schemas.ErrorSchema{
				Code:    "ErrCreateUser",
				Message: "Internal server error: Could not create a user",
			},
		})
		return
	}

	json.NewEncoder(res).Encode(response)
}

func (handler *AuthHandler) Login(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	log.Println("/auth/login")

	var body schemas.LoginRequest
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
