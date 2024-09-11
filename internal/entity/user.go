package entity

import "github.com/google/uuid"

type User struct {
	ID            uuid.UUID
	LastIPAddress string
	Tokens        *Tokens
}

func NewUser(id uuid.UUID, lastIPAddress string) *User {
	return &User{
		ID:            id,
		LastIPAddress: lastIPAddress,
	}
}
