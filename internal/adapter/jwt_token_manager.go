package adapter

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"time"

	"Ang2Tea/medods-test/internal/usecase"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	oneYear = 24 * time.Hour
)

type CustomClaims struct {
	jwt.RegisteredClaims

	IPAddress string `json:"ip_address"`
}

var _ usecase.ITokenManager = (*jwtTokenManager)(nil)

type jwtTokenManager struct {
	secretKey string
}

func NewJWTTokenManager(secret string) *jwtTokenManager {
	return &jwtTokenManager{
		secretKey: secret,
	}
}

func (j *jwtTokenManager) Generate(ctx context.Context, userInfo usecase.UserRequest) (*usecase.Tokens, error) {
	accessToken, err := j.genAccessToken(userInfo.UserID, userInfo.IPAddress)
	if err != nil {
		return nil, err
	}

	refreshToken, err := j.genRefreshToken(accessToken)
	if err != nil {
		return nil, err
	}

	return &usecase.Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func (j *jwtTokenManager) genAccessToken(userID uuid.UUID, ipAddress string) (string, error) {
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(oneYear)),
		},
		IPAddress: ipAddress,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString(j.secretKey)
}

func (j *jwtTokenManager) genRefreshToken(accessToken string) (string, error) {
	rawKey := append([]byte(accessToken), []byte(j.secretKey)...)
	s := sha1.Sum(rawKey)

	return base64.StdEncoding.EncodeToString(s[:]), nil
}
