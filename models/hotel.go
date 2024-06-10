package models

import (
	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	Name      string          `json:"name"`
	Locations []HotelLocation `gorm:"foreignKey:HotelID" json:"locations"`
}
