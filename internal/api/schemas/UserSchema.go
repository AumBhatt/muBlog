package schemas

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	MailId   string `json:"mailId" validate:"required,email"`
}

type CreateUserResponse struct {
	Id 		 *string `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
	*ErrorSchema
}

