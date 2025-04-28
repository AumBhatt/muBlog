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

type GetByPostIdResponse struct {
	PostId    string `json:"postId"`
	AuthorId  string `json:"authorId"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"createdAt"`
	EditedAt  int64  `json:"editedAt"`
	GetReactionsCountByPostIdResponse
}

type GetByUserIdRequest struct{}

type GetByUserIdResponse struct{}

type AddReactionRequest struct {
	PostId string `json:"postId" validate:"required,uuid"`
	UserId string `json:"userId" validate:"required,uuid"`
	Type   string `json:"type" validate:"required"`
}

type AddReactionResponse struct {
	Reactions []map[string]any `json:"reactions"`
}

type GetReactionsByPostIdRequest struct{}

type GetReactionsByPostIdResponse struct {
	Reactions []map[string]string `json:"reactions"`
}

type GetReactionsCountByPostIdRequest struct{}

type GetReactionsCountByPostIdResponse struct {
	Reactions []map[string]any `json:"reactions"`
}
