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

	response, err := handler.userService.Unfollow(body)
	if err != nil {
		log.Println("UserHandler.Unfollow err:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(res).Encode(response)
	if err != nil {
		log.Println("UserHandler.Unfollow err:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *UserHandler) GetFollowersById(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	userId := ps.ByName("userId")
	if userId == "" {
		log.Println("UserHandler.GetFollowersById err: userId missing")
		res.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(res).Encode(&schemas.ErrorSchema{
			Code:    "MissingUserId",
			Message: "Error: userId is missing in path params.",
		})
		if err != nil {
			log.Println("UserHandler.GetFollowersById err:", err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	data, err := handler.userService.GetFollowersById(userId)
	if err != nil {
		log.Println("UserHandler.GetFollowersById err:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(res).Encode(data)
	if err != nil {
		log.Println("UserHandler.GetFollowersById err:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *UserHandler) GetFollowingById(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	followersId := ps.ByName("followersId")
	if followersId == "" {
		log.Println("UserHandler.GetFollowingById err: followersId missing")
		res.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(res).Encode(&schemas.ErrorSchema{
			Code:    "MissingFollowersId",
			Message: "Error: followersId is missing in path params.",
		})
		if err != nil {
			log.Println("UserHandler.GetFollowingById err:", err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	data, err := handler.userService.GetFollowingById(followersId)
	if err != nil {
		log.Println("UserHandler.GetFollowingById err:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(res).Encode(data)
	if err != nil {
		log.Println("UserHandler.GetFollowingById err:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}
