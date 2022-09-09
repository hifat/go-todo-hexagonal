package repository

type Register struct {
	Username string `db:"username"`
	Password string `db:"password"`
	Name     string `db:"name"`
}

type Login struct {
	Username string `db:"username"`
	Password string `db:"password"`
}

type Auth struct {
	Token string `db:"token"`
	User
}

type AuthRepository interface {
	Register(register Register) (*Auth, error)
	Login(login Login) (*Auth, error)
	Me(token string) (*Auth, error)
}
