package service

import "time"

type RenewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RenewAccessTokenResponse struct {
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}

type UserDevice struct {
	UserAgent string `json:"user_agent"`
	ClientIP  string `json:"client_ip"`
}

type TokenService interface {
	RenewAccessToken(renewAccessTokenReq RenewAccessTokenRequest, userDevice UserDevice) (*RenewAccessTokenResponse, error)
}
