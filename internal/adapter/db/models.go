package db

import "Ang2Tea/medods-test/internal/entity"

type User struct {
	UUIDModel

	LastIPAddress string
	RefreshToken  string
}

func (u *User) ToDomain() entity.User {
	return entity.User{
		ID:            u.ID,
		LastIPAddress: u.LastIPAddress,
		RefreshToken:  u.RefreshToken,
	}
}

func (u *User) FromDomain(user entity.User) User {
	return User{
		UUIDModel: UUIDModel{
			ID: user.ID,
		},
		LastIPAddress: user.LastIPAddress,
		RefreshToken:  user.RefreshToken,
	}
}
