package models

import "time"

type TrainingRegistration struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	TrainingSessionID uint      `json:"training_session_id"`
	UserID            uint      `json:"user_id"`
	RegistrationDate  time.Time `gorm:"autoCreateTime" json:"registration_date"`
	Status            string    `gorm:"size:50" json:"status"` 
}
