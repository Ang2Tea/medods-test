package entity

import "github.com/google/uuid"

type User struct {
	ID            uuid.UUID
	LastIPAddress string
	AccessToken   *string
	RefreshToken  string
}

func NewUser(id uuid.UUID, lastIPAddress string) *User {
	return &User{
		ID:            id,
		LastIPAddress: lastIPAddress,
	}
}

func (u *User) SetTokens(accessToken, refreshToken string) {
	u.AccessToken = &accessToken
	u.RefreshToken = refreshToken
}
