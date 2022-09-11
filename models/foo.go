package models

type Foo struct {
	ID        int `gorm:"primaryKey"`
	Title     string
	Price     float64
	CreatedAt string
	UpdatedAt string
}
