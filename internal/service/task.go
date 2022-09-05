package service

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
	UserID uint   `json:"user_id"`
	Detail string `json:"detail"`
	Done   bool   `json:"done"`
}

type EditTask struct {
	Detail string `json:"detail"`
	Done   bool   `json:"done"`
}

type TaskService interface {
	Get() ([]Task, error)
	Create(task NewTask) (*Task, error)
	Show(id uint) (*Task, error)
	Update(id uint, task EditTask) (*Task, error)
	Delete(id uint) error
}
