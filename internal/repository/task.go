package repository

import (
	"time"
)

type Task struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	Detail    string    `db:"detail"`
	Done      bool      `db:"done"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type NewTask struct {
	UserID string `db:"user_id"`
	Detail string `db:"detail"`
	Done   bool   `db:"done"`
}

type EditTask struct {
	Detail string `db:"detail"`
	Done   bool   `db:"done"`
}

type TaskRepository interface {
	Get() ([]Task, error)
	Create(task NewTask) (*Task, error)
	Show(id string) (*Task, error)
	ToggleDone(id string) (*Task, error)
	Update(id string, task EditTask) (*Task, error)
	Delete(id string) error
}
