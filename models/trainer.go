package models

import "time"

type Trainer struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TrainerType string    `gorm:"size:50" json:"trainer_type"`
	FullName    string    `gorm:"size:255" json:"full_name"`
	Age         int       `json:"age"`
	Height      float64   `json:"height"`
	Weight      float64   `json:"weight"`
	Country     string    `gorm:"size:100" json:"country"`
	About       string    `json:"about"`
	ImageURL    string    `gorm:"size:255" json:"image_url"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
