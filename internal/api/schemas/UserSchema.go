package schemas

type GetUserByIdRequest struct {
	Id string `json:"id" validate:"required,min=3,max=20"`
}

type GetUserByIdResponse struct {
	Id          string `json:"id,omitempty"`
	Username    string `json:"username,omitempty"`
	MailId      string `json:"mailId,omitempty"`
	ActiveSince int64  `json:"activeSince,omitempty"`
	*ErrorSchema
}
