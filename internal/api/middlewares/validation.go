package middlewares

import (
	"encoding/json"
	"fmt"
	"muBlog/internal/api/schemas"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func ValidateRequest[schema any](res http.ResponseWriter, req *http.Request, params httprouter.Params) func(httprouter.Handle) httprouter.Handle {

	// execute this wrapper func each time the endpoint is been hit
	return func(handler httprouter.Handle) httprouter.Handle {

		var body schema

		json.NewDecoder(req.Body).Decode(body)
		validate := validator.New()
		err := validate.Struct(body)

		var response schemas.ValidationErrorSchema
		response.Code = "ValidationErrors"

		if err != nil {
			errors := err.(validator.ValidationErrors)
			for _, err := range errors {
				response.Errors = append(response.Errors, schemas.ErrorSchema{
					Code:    fmt.Sprintf("Invalid%s", err.ActualTag()),
					Message: err.Error(),
				})
			}

			json.NewEncoder(res).Encode(response)
			return nil
		}

		return handler
	}
}
