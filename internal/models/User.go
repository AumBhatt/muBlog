package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id          string
	Username    string
	MailId      string
	ActiveSince int64
}

func NewUser(username string, mailId string) *User {
	return &User{
		Id:          fmt.Sprintf("U-%s", uuid.New().String()),
		Username:    username,
		MailId:      mailId,
		ActiveSince: time.Time.UnixMilli(time.Now()),
	}
}
