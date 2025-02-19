package models

import "time"

type Payment struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	UserID        uint      `json:"user_id"`
	OrderID       uint      `json:"order_id"`
	Amount        float64   `gorm:"type:decimal(10,2)" json:"amount"`
	PaymentDate   time.Time `gorm:"autoCreateTime" json:"payment_date"`
	PaymentMethod string    `gorm:"size:50" json:"payment_method"`
}
