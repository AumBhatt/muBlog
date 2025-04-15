package schemas

type ErrorSchema struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ValidationErrorSchema struct {
	Code   string        `json:"code"`
	Errors []ErrorSchema `json:"errors"`
}
