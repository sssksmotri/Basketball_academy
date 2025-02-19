package models

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:255" json:"name"`
	Description string    `json:"description"`
	Price       float64   `gorm:"type:decimal(10,2)" json:"price"`
	Stock       int       `json:"stock"`
	ImageURL    string    `gorm:"size:255" json:"image_url"`
	Category    string    `gorm:"size:100" json:"category"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
