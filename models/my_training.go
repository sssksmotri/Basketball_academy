package models

import "time"

type MyTraining struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	TrainingDate time.Time `json:"training_date"`
	TrainingTime time.Time `json:"training_time"`
	NumPlaces    int       `json:"num_places"`
	Address      string    `gorm:"size:255" json:"address"`
	TrainerID    uint      `json:"trainer_id"`
	Status       string    `gorm:"size:50" json:"status"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
