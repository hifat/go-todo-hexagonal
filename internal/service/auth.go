package service

import "time"

type Register struct {
	Username  string `validate:"required,max=100" json:"username"`
	Password  string `validate:"required,max=100" json:"password"`
	Name      string `validate:"required,max=100" json:"name"`
	UserAgent string `json:"user_agent"`
	ClientIP  string `json:"client_ip"`
}

type Login struct {
	Username  string `validate:"required,max=100" json:"username"`
	Password  string `validate:"required,max=150" json:"password"`
	UserAgent string `json:"user_agent"`
	ClientIP  string `json:"client_ip"`
}

type Auth struct {
	User                  User      `json:"user"`
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
