package handler

import "time"

type Task struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Detail    string    `json:"detail"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewTask struct {
	Detail string `json:"detail"`
}

type EditTask struct {
	Detail string `json:"detail"`
	Done   bool   `json:"done"`
}

type TaskHandler interface {
	Get() ([]Task, error)
	Create(task NewTask) (*Task, error)
	Show(task Task, id uint) (*Task, error)
	Update(id uint, task EditTask) (*Task, error)
	Delete(id uint) error
}
