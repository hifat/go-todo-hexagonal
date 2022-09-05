package migration

import "time"

type Task struct {
	ID        uint
	UserID    uint
	Detail    string `gorm:"type:varchar(255)"`
	Done      bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
