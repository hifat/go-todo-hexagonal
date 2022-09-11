package handler

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
