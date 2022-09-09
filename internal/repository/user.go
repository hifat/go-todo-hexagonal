package repository

type User struct {
	ID        string `db:"id"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	Name      string `db:"name"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
