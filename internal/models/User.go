package models

type User struct {
	Id          string
	Username    string
	Email       string
	ActiveSince int64
	Password    string
}

type Follow struct {
	Id         string
	UserId     string
	FollowerId string
}
