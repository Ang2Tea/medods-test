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
	events       IAuthEvent
}

func NewAuthUsecase(
	userDB IUserStorage,
	tokenManager ITokenManager,
	events IAuthEvent,
) *authUsecase {
	return &authUsecase{
		userDB:       userDB,
		tokenManager: tokenManager,
		events:       events,
	}
}

func (a *authUsecase) Register(ctx context.Context, request UserRequest) (*Tokens, error) {
	user, err := a.userDB.GetByID(ctx, request.UserID)
	if err != nil && !errors.Is(err, entity.ErrUserNotFound) {
		return nil, err
	}

	if user == nil {
		user = entity.NewUser(request.UserID, request.IPAddress)

		_, err = a.userDB.Create(ctx, *user)
		if err != nil {
			return nil, err
		}
	}

	return a.generateToken(ctx, user, request)
}

func (a *authUsecase) Refresh(ctx context.Context, request RefreshRequest) (*Tokens, error) {
	user, err := a.userDB.GetByRefreshToken(ctx, request.RefreshToken)
	if err != nil {
		return nil, err
	}

	generateRequest := UserRequest{
		UserID:    user.ID,
		IPAddress: user.LastIPAddress,
	}

	if user.LastIPAddress != request.IPAddress && a.events != nil {
		a.events.IPAddressChanged(ctx, user.LastIPAddress, request.IPAddress)
	}

	return a.generateToken(ctx, user, generateRequest)
}

func (a *authUsecase) generateToken(ctx context.Context, user *entity.User, request UserRequest) (*Tokens, error) {
	token, err := a.tokenManager.Generate(ctx, request)
	if err != nil {
		return nil, err
	}

	user.SetTokens(token.Access, token.Refresh)
	user.LastIPAddress = request.IPAddress

	_, err = a.userDB.Update(ctx, *user)
	if err != nil {
		return nil, err
	}

	return token, nil
}
