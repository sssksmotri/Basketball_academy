package models

import "time"

type StaticPage struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Slug      string    `gorm:"size:100;unique" json:"slug"`
	Title     string    `gorm:"size:255" json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
