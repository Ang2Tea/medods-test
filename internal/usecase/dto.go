package usecase

import (
	"github.com/google/uuid"
)

type Tokens struct {
	Access  string
	Refresh string
}

type UserRequest struct {
	UserID    uuid.UUID
	IPAddress string
}

type RefreshRequest struct {
	RefreshToken string
	IPAddress    string
}
