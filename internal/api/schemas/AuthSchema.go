package schemas

type SignupRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	MailId   string `json:"mailId" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=40"`
}

type SignupResponse struct {
	Id       *string `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
	*ErrorSchema
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLoginResponse struct {
	Id    *string `json:"id,omitempty"`
	Token *string `json:"token,omitempty"`
	*ErrorSchema
}
