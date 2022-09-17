package Models

import "time"

type Product struct {
	ID          int `gorm:"primaryKey"`
	Title       string
	Description string
	Price       string
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
