package models

import "gorm.io/gorm"

type Amenity struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
}
