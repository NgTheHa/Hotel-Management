package models

import "gorm.io/gorm"

type Hotel struct {
	gorm.Model
	Name       string     `json:"name"`
	Address    string     `json:"address"`
	RoomCount  int        `json:"room_count"`
	Facilities []Facility `gorm:"many2many:hotel_facilities;" json:"facilities"`
}
