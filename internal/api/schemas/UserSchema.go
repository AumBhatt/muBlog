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

type FollowRequest struct {
	UserId     string `json:"userId" validate:"required,uuid"`
	FollowerId string `json:"followerId" validate:"required,uuid"`
}

type FollowResponse struct {
	FollowId string `json:"followId,omitempty"`
	*ErrorSchema
}

type UnfollowRequest struct {
	UserId     string `json:"userId" validate:"required,uuid"`
	FollowerId string `json:"followerId" validate:"required,uuid"`
}

type UnfollowResponse struct {
	Status string
	*ErrorSchema
}
