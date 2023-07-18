package entity

import (
	"time"
)

type User struct {
	ID        int64 `gorm:"primaryKey,omitempty"`
	Name      string
	Age       int
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
