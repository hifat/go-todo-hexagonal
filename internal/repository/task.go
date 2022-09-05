package repository

import (
	"time"
)

type Task struct {
	ID        uint      `db:"id"`
	UserID    uint      `db:"user_id"`
	Detail    string    `db:"detail"`
	Done      bool      `db:"done"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type NewTask struct {
	UserID uint   `db:"user_id"`
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
	Show(id uint) (*Task, error)
	Update(id uint, task EditTask) (*Task, error)
	Delete(id uint) error
}
