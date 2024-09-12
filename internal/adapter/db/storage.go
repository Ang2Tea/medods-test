package db

import (
	"context"
	"errors"

	"Ang2Tea/medods-test/internal/entity"
	"Ang2Tea/medods-test/internal/usecase"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ usecase.IUserStorage = (*userStorage)(nil)

type userStorage struct {
	db           *gorm.DB
	accessTokens map[uuid.UUID]string
}

func NewUserStorage(db *gorm.DB) *userStorage {
	return &userStorage{
		db:           db,
		accessTokens: make(map[uuid.UUID]string, 32),
	}
}

func (u *userStorage) Create(ctx context.Context, user entity.User) (*uuid.UUID, error) {
	tx := u.getTX(ctx)

	dbUser := new(User).FromDomain(user)

	tx = tx.Create(&dbUser)
	if err := returnNormalizeErr(tx.Error); err != nil {
		return nil, err
	}

	if user.AccessToken != nil {
		u.accessTokens[dbUser.ID] = *user.AccessToken
	}

	return &dbUser.ID, nil
}

func (u *userStorage) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	tx := u.getTX(ctx)

	var dbUser User

	tx = tx.First(&dbUser, id)
	if err := returnNormalizeErr(tx.Error); err != nil {
		return nil, err
	}

	result := dbUser.ToDomain()

	if userAccessToken, ok := u.accessTokens[id]; ok {
		result.AccessToken = &userAccessToken
	}

	return &result, nil
}

func (u *userStorage) GetByRefreshToken(ctx context.Context, refreshToken string) (*entity.User, error) {
	tx := u.getTX(ctx)

	var dbUser User

	tx = tx.First(&dbUser, &User{RefreshToken: refreshToken})
	if err := returnNormalizeErr(tx.Error); err != nil {
		return nil, err
	}

	result := dbUser.ToDomain()

	if userAccessToken, ok := u.accessTokens[result.ID]; ok {
		result.AccessToken = &userAccessToken
	}

	return &result, nil
}

func (u *userStorage) Update(ctx context.Context, user entity.User) (*uuid.UUID, error) {
	tx := u.getTX(ctx)

	dbUser := new(User).FromDomain(user)

	tx = tx.Where(User{UUIDModel: UUIDModel{ID: user.ID}}).
		Updates(dbUser)

	if err := returnNormalizeErr(tx.Error); err != nil {
		return nil, err
	}

	u.accessTokens[user.ID] = *user.AccessToken

	return &user.ID, nil
}

func (u *userStorage) getTX(ctx context.Context) *gorm.DB {
	return u.db.WithContext(ctx).Model(new(User))
}

func returnNormalizeErr(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.ErrUserNotFound
	}

	return err
}
