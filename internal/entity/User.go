package entity

import (
	"time"
)

type User struct {
	ID        int64 `gorm:"primaryKey,omitempty,auto_increment"`
	Name      string
	Age       int
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
