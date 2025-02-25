package models

import "time"

// Модель пользователя
type User struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	FirstName     string    `gorm:"size:100" json:"first_name"`
	LastName      string    `gorm:"size:100" json:"last_name"`
	Email         string    `gorm:"size:255;unique" json:"email"`
	Phone         string    `gorm:"size:20;uniqueIndex" json:"phone"`
	PasswordHash  string    `gorm:"size:255" json:"password_hash"`
	Role          string    `gorm:"size:20" json:"role"`
	WalletBalance float64   `gorm:"type:decimal(10,2);default:0" json:"wallet_balance"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
