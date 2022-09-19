package Models

import (
	"time"
)

type Product struct {
	ID           int `gorm:"primaryKey"`
	Title        string
	Description  string
	Price        string
	IsDeleted    bool           `gorm:"default:false"`
	ProductImage []ProductImage `gorm:"foreignKey:ProductID;references:ID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ProductImage struct {
	ID        int `gorm:"primaryKey"`
	Src       string
	ProductID int
	CreatedAt time.Time
	UpdatedAt time.Time
}
