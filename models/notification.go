package models

import "time"

type Notification struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	Title     string    `gorm:"size:255" json:"title"`
	Message   string    `json:"message"`
	Type      string    `gorm:"size:50" json:"type"` // 'chat', 'news', 'training', 'payment'
	IsRead    bool      `gorm:"default:false" json:"is_read"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
