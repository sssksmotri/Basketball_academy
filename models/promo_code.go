package models

import "time"

type PromoCode struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	Code               string    `gorm:"size:50;unique" json:"code"`
	DiscountPercentage float64   `gorm:"type:decimal(5,2)" json:"discount_percentage"`
	Description        string    `json:"description"`
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
	UsageLimit         int       `json:"usage_limit"`
	TimesUsed          int       `gorm:"default:0" json:"times_used"`
	Active             bool      `gorm:"default:true" json:"active"`
}
