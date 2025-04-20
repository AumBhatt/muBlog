package schemas

type GetUserByIdRequest struct {
	Id string `json:"id" validate:"required,uuid"`
}

type GetUserByIdResponse struct {
	Id          string `json:"id,omitempty"`
	Username    string `json:"username,omitempty"`
	Email       string `json:"email,omitempty"`
	ActiveSince int64  `json:"activeSince,omitempty"`
	*ErrorSchema
}
