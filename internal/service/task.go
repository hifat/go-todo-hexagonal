package service

import "time"

type Task struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Detail    string    `json:"detail"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewTask struct {
	UserID string `json:"user_id"`
	Detail string `json:"detail"`
	Done   bool   `json:"done"`
}

type EditTask struct {
	Detail string `json:"detail"`
	Done   bool   `json:"done"`
}

type TaskService interface {
	Get(userID string) ([]Task, error)
	Create(task NewTask) (*Task, error)
	Show(id string) (*Task, error)
	ToggleDone(id string) (*Task, error)
	Update(id string, task EditTask) (*Task, error)
	Delete(id string) error
}
