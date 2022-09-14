package migration

import (
	"time"
)

type Task struct {
	ID        uint
	UserID    uint      `gorm:"type:varchar(20)not null"`
	Detail    string    `gorm:"type:varchar(255); not null"`
	Done      bool      `gorm:"default:false; not null"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
