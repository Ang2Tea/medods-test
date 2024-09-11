package usecase

import (
	"context"
	"errors"

	"Ang2Tea/medods-test/internal/entity"
)

var _ IAuth = (*authUsecase)(nil)

type authUsecase struct {
	userDB       IUserStorage
	tokenManager ITokenManager
}

func NewAuthUsecase(userDB IUserStorage, tokenManager ITokenManager) IAuth {
	return &authUsecase{
		userDB:       userDB,
		tokenManager: tokenManager,
	}
}

func (a *authUsecase) Register(ctx context.Context, request UserRequest) (*User, error) {
	user, err := a.userDB.GetByID(ctx, request.UserID)
	if err != nil && !errors.Is(err, entity.ErrUserNotFound) {
		return nil, err
	}

	if user == nil {
		user = entity.NewUser(request.UserID, request.IPAddress)

		token, err := a.tokenManager.Generate(ctx, request)
		if err != nil {
			return nil, err
		}

		domainToken := new(Tokens).ToDomain(*token)
		user.Tokens = &domainToken

		_, err = a.userDB.Create(ctx, *user)
		if err != nil {
			return nil, err
		}
	}

	result := new(User).FromDomain(*user)

	return &result, nil
}

func (a *authUsecase) Refresh(ctx context.Context, request RefreshRequest) (*User, error) {
	user, err := a.userDB.GetByRefreshToken(ctx, request.RefreshToken)
	if err != nil {
		return nil, err
	}

	generateRequest := UserRequest{
		UserID:    user.ID,
		IPAddress: user.LastIPAddress,
	}

	token, err := a.tokenManager.Generate(ctx, generateRequest)
	if err != nil {
		return nil, err
	}

	domainToken := new(Tokens).ToDomain(*token)
	user.Tokens = &domainToken

	_, err = a.userDB.Update(ctx, *user)
	if err != nil {
		return nil, err
	}

	result := new(User).FromDomain(*user)

	return &result, nil
}
