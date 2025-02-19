package models

import "time"

type TrainingSession struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	TrainerID     uint      `json:"trainer_id"` 
	TrainingType  string    `gorm:"size:100" json:"training_type"`
	TrainingDate  time.Time `json:"training_date"`
	StartTime     string    `json:"start_time"`
	EndTime       string    `json:"end_time"`
	Location      string    `gorm:"size:255" json:"location"`
	Latitude      float64   `gorm:"type:decimal(9,6)" json:"latitude"`
	Longitude     float64   `gorm:"type:decimal(9,6)" json:"longitude"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
