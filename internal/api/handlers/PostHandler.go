package handlers

import (
	"encoding/json"
	"log"
	"muBlog/internal/api/schemas"
	"muBlog/internal/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PostHandler struct {
	postService *services.PostService
}

func NewPostHandler(postService *services.PostService) *PostHandler {
	return &PostHandler{postService}
}

func (handler *PostHandler) Create(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	var body schemas.CreatePostRequest

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Println("PostHandler.Create:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := handler.postService.CreatePost(body)
	if err != nil {
		log.Println("PostHandler.Create:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(res).Encode(response)
	if err != nil {
		log.Println("PostHandler.Create:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *PostHandler) React(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

}
