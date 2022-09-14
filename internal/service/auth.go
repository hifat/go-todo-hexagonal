package service

import "time"

type Register struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	UserAgent string `json:"user_agent"`
	ClientIP  string `json:"client_ip"`
}

type Login struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	UserAgent string `json:"user_agent"`
	ClientIP  string `json:"client_ip"`
}

type Auth struct {
	User
	SessionID             string    `json:"session_id"`
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}

type AuthService interface {
	Register(register Register) (*Auth, error)
	Login(login Login) (*Auth, error)
	Me(token string) (*Auth, error)
}
