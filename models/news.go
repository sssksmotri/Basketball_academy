package models

import "time"

// Модель новостей
type News struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:255" json:"title"`
	Content     string    `json:"content"`
	ImageURL    string    `gorm:"size:255" json:"image_url"`
	NotifyUsers bool      `gorm:"default:false" json:"notify_users"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
