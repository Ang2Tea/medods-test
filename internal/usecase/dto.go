package usecase

import (
	"Ang2Tea/medods-test/internal/entity"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID
	LastIPAddress string
	Tokens        *Tokens
}

func (*User) ToDomain(u User) entity.User {
	result := entity.User{
		ID:            u.ID,
		LastIPAddress: u.LastIPAddress,
	}

	if u.Tokens != nil {
		token := new(Tokens).ToDomain(*u.Tokens)
		result.Tokens = &token
	}

	return result
}

func (*User) FromDomain(u entity.User) User {
	result := User{
		ID:            u.ID,
		LastIPAddress: u.LastIPAddress,
	}

	if u.Tokens != nil {
		token := new(Tokens).FromDomain(*u.Tokens)
		result.Tokens = &token
	}

	return result
}

type Tokens struct {
	Access  string
	Refresh string
}

func (*Tokens) ToDomain(t Tokens) entity.Tokens {
	result := entity.Tokens{
		Access:  t.Access,
		Refresh: t.Refresh,
	}

	return result
}

func (*Tokens) FromDomain(t entity.Tokens) Tokens {
	result := Tokens{
		Access:  t.Access,
		Refresh: t.Refresh,
	}

	return result
}

type UserRequest struct {
	UserID    uuid.UUID
	IPAddress string
}

type RefreshRequest struct {
	RefreshToken string
	IPAddress    string
}
