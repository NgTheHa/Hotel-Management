package models

import (
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	BookingID   uint    `json:"booking_id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
