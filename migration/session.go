package migration

import (
	"time"
)

type Session struct {
	ID           string    `gorm:"type:varchar(255);primary_key"`
	Username     string    `gorm:"type:varchar(100);not null"`
	RefreshToken string    `gorm:"type:text;not null"`
	UserAgent    string    `gorm:"type:text;not null"`
	ClientIP     string    `gorm:"type:varchar(100);not null"`
	IsBlocked    bool      `gorm:"default:false;not null"`
	ExpiresAt    time.Time `gorm:"not null"`
	CreatedAt    time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
