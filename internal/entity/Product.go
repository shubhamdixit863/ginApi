package entity

import "time"

// Product

type Product struct {
	ID           int64 `gorm:"primaryKey,omitempty"`
	Name         string
	ProductImage string
	Description  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
