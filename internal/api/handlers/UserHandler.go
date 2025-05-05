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

func (handler *UserHandler) Follow(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	var body *schemas.FollowRequest
	err := json.NewDecoder(req.Body).Decode(body)
	if err != nil {
		log.Println("UserHandler.Follow err:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := handler.userService.Follow(body)
	if err != nil {
		log.Println("UserHandler.Follow err:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(res).Encode(response)
	if err != nil {
		log.Println("UserHandler.Follow err:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *UserHandler) Unfollow(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	var body *schemas.UnfollowRequest
	err := json.NewDecoder(req.Body).Decode(body)
	if err != nil {
		log.Println("UserHandler.Unfollow err:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// response, err := handler.
}
