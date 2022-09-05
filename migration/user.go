package migration

import "time"

type User struct {
	ID        uint
	Username  string `gorm:"type:varchar(100)"`
	Password  string `gorm:"type:varchar(255)"`
	Name      string `gorm:"type:varchar(100)"`
	Available bool   `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
