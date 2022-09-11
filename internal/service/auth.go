package service

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Auth struct {
	AccessToken string `json:"access_token"`
	User
}

type AuthService interface {
	Register(register Register) (*Auth, error)
	Login(login Login) (*Auth, error)
	Me(token string) (*Auth, error)
}
