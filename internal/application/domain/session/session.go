package dmsession

import "google.golang.org/genproto/googleapis/type/datetime"

type Session struct {
	Id                    string
	UserId                string
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  *datetime.DateTime
	RefreshTokenExpiresAt *datetime.DateTime
}
