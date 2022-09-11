package migration

import "time"

type User struct {
	ID        uint
	Username  string    `gorm:"type:varchar(100);unique;not null"`
	Password  string    `gorm:"type:varchar(150);not null"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Available bool      `gorm:"default:true;not null"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
