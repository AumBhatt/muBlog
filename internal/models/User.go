package models

type User struct {
	Id          string
	Username    string
	MailId      string
	ActiveSince int64
	Password    string
}
