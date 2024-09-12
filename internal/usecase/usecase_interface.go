package usecase

import (
	"context"

	"github.com/google/uuid"

	"Ang2Tea/medods-test/internal/entity"
)

type IUserStorage interface {
	Create(ctx context.Context, user entity.User) (*uuid.UUID, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetByRefreshToken(ctx context.Context, refreshToken string) (*entity.User, error)
	Update(ctx context.Context, user entity.User) (*uuid.UUID, error)
}

type ITokenManager interface {
	Generate(ctx context.Context, userInfo UserRequest) (*Tokens, error)
}

type IAuthEvent interface {
	IPAddressChanged(ctx context.Context, oldIPAddress, newIPAddress string)
}

type IAuth interface {
	Register(ctx context.Context, request UserRequest) (*Tokens, error)
	Refresh(ctx context.Context, request RefreshRequest) (*Tokens, error)
}
