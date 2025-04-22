package schemas

type CreatePostRequest struct {
	Content  string `json:"content" validate:"required"`
	AuthorId string `json:"authorId" validate:"required"`
}

type CreatePostResponse struct {
	Id string `json:"postId"`
}

type EditRequest struct{}

type EditResponse struct{}

type DeleteRequest struct{}

type DeleteResponse struct{}

type GetByIdRequest struct{}

type GetByIdResponse struct{}

type GetByUserIdRequest struct{}

type GetByUserIdResponse struct{}

type AddReactionRequest struct {
	PostId       string `json:"postId" validate:"required,uuid"`
	UserId       string `json:"userId" validate:"required,uuid"`
	ReactionType string `json:"reactionType" validate:"required"`
}

type AddReactionResponse struct{}

type GetReactionCountsRequest struct{}

type GetReactionCountsResponse struct{}

type ReactionsRequest struct{}

type ReactionsResponse struct{}
