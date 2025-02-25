package models

import "time"

type Club struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:255" json:"name"`
	Address     string    `gorm:"size:255" json:"address"`
	Coordinates string    `gorm:"size:255" json:"coordinates"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
