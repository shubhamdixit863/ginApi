package entity

import (
	"time"
)

type User struct {
	ID        int64 `gorm:"primaryKey,omitempty"`
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
