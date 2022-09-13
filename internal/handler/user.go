package handler

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type EditUser struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}
