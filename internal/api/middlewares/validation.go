package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"muBlog/internal/api/schemas"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func ValidateRequest[schema any](next httprouter.Handle) func(http.ResponseWriter, *http.Request, httprouter.Params) {

	// execute this wrapper func each time the endpoint is been hit
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		var body schema
		err = json.Unmarshal(bodyBytes, &body)
		if err != nil {
			log.Println(err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		validate := validator.New()
		err = validate.Struct(body)

		var response schemas.ValidationErrorSchema
		response.Code = "ValidationErrors"

		if err != nil {
			errors := err.(validator.ValidationErrors)
			for _, err := range errors {
				response.Errors = append(response.Errors, schemas.ErrorSchema{
					Code:    fmt.Sprintf("Invalid%s", err.Field()),
					Message: fmt.Sprintf("%s is %s", err.Field(), err.Tag()),
				})
			}

			json.NewEncoder(res).Encode(response)
			return
		}

		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		next(res, req, params)
	}
}
