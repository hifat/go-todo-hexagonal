package service

import "time"

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EditUser struct {
	Username string `validate:"required,max=100" json:"username"`
	Name     string `validate:"required,max=100" json:"name"`
}

type UserService interface {
	Update(id string, editUser EditUser) (*User, error)
}
