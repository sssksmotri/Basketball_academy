package models

import "time"

type Chat struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	TrainerID uint      `json:"trainer_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
