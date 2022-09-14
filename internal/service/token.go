package service

import "time"

type RenewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RenewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

type TokenService interface {
	RenewAccessToken(renewAccessTokenReq RenewAccessTokenRequest) (*RenewAccessTokenResponse, error)
}
