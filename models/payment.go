package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	BookingID uint    `json:"booking_id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
}
