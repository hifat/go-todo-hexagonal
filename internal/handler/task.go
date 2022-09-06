package handler

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
	Detail string `json:"detail"`
}

type EditTask struct {
	Detail string `json:"detail"`
	Done   bool   `json:"done"`
}
