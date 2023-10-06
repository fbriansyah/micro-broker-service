package dmsession

import "time"

type Session struct {
	Id                    string
	UserId                string
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  time.Time
	RefreshTokenExpiresAt time.Time
}
