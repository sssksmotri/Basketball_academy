package models

import "time"

type Order struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	UserID        uint      `json:"user_id"`
	TotalAmount   float64   `gorm:"type:decimal(10,2)" json:"total_amount"`
	Status        string    `gorm:"size:50" json:"status"` 
	PaymentMethod string    `gorm:"size:50" json:"payment_method"`
	OrderDate     time.Time `gorm:"autoCreateTime" json:"order_date"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
