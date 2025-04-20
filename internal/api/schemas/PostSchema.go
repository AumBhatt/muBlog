package schemas

type CreatePostRequest struct {
	Content  string `json:"content" validate:"required,uuid"`
	AuthorId string `json:"authorId" validate:"required"`
}

type CreatePostResponse struct{}

type EditRequest struct{}

type EditResponse struct{}

type DeleteRequest struct{}

type DeleteResponse struct{}

type GetByIdRequest struct{}

type GetByIdResponse struct{}

type GetByUserIdRequest struct{}

type GetByUserIdResponse struct{}

type ReactRequest struct{}

type ReactResponse struct{}

type ReactionsRequest struct{}

type ReactionsResponse struct{}
