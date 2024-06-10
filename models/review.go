package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Rating     float64 `json:"rating"`
	Comment    string  `json:"comment"`
	CustomerID uint
	HotelID    uint
}
