package handler

import "Ang2Tea/medods-test/internal/usecase"

type TokensResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func toFiber(t usecase.Tokens) *TokensResponse {
	return &TokensResponse{
		AccessToken:  t.Access,
		RefreshToken: t.Refresh,
	}
}