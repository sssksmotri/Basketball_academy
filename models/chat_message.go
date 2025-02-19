package models

import "time"

type ChatMessage struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ChatID    uint      `json:"chat_id"`
	SenderID  uint      `json:"sender_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
