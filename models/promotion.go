package models

import "time"

type Promotion struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:255" json:"title"`
	Description string    `json:"description"`
	Discount    float64   `gorm:"type:decimal(5,2)" json:"discount"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	ImageURL    string    `gorm:"size:255" json:"image_url"`
	NotifyUsers bool      `gorm:"default:false" json:"notify_users"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
