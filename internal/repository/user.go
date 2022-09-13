package repository

import "time"

type User struct {
	ID        string    `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type EditUser struct {
	Username string `db:"detail"`
	Name     string `db:"name"`
}

type UserRepository interface {
	Update(id string, editUser EditUser) (*User, error)
}
