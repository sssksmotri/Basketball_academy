package models

import "time"

type AcademyStaff struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:255" json:"name"`
	Position    string    `gorm:"size:100" json:"position"`
	Bio         string    `json:"bio"`
	ContactInfo string    `gorm:"size:255" json:"contact_info"`
	ImageURL    string    `gorm:"size:255" json:"image_url"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
