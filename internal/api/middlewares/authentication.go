package middlewares

import (
	"encoding/json"
	"log"
	"muBlog/internal/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Authentication(authService *services.AuthService, next httprouter.Handle) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		token := req.Header.Get("Authorization")
		if token == "" {
			res.WriteHeader(http.StatusUnauthorized)
			return
		}

		verificationFailed, err := authService.VerifyToken(token)
		if err != nil {
			log.Println("Authentication err:", err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		if verificationFailed != nil {
			res.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(res).Encode(verificationFailed)
			if err != nil {
				res.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}

		next(res, req, ps)

	}
}
