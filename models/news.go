package models

import "time"

type News struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Category    string    `gorm:"size:255" json:"category"`
	Title       string    `gorm:"size:255" json:"title"`
	NewsDate    time.Time `json:"news_date"`
	StartTime   time.Time `json:"start_time"`
	Address     string    `gorm:"size:255" json:"address"`
	PublishedAt time.Time `json:"published_at"`
	Description string    `json:"description"`
	VosJdut     *string   `json:"vos_jdut,omitempty"`
	ImageURL    string    `gorm:"size:255" json:"image_url"`
	NotifyUsers bool      `gorm:"default:false" json:"notify_users"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
