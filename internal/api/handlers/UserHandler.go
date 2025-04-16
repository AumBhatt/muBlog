package handlers

import (
	"encoding/json"
	"log"
	"muBlog/internal/api/schemas"
	"muBlog/internal/services"
	"net/http"

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
		Email:       user.Email,
		ActiveSince: user.ActiveSince,
	})
}
