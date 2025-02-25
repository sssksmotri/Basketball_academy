package models

import "time"

type SupportTicket struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `json:"user_id"`
	Subject     string    `gorm:"size:255" json:"subject"`
	Description string    `json:"description"`
	Status      string    `gorm:"size:50" json:"status"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
