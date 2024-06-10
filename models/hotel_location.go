package models

import "gorm.io/gorm"

type HotelLocation struct {
	gorm.Model
	Name    string  `json:"name"`
	Address Address `gorm:"embedded" json:"address"`
	Rooms   []Room  `gorm:"foreignKey:LocationID" json:"rooms"`
	HotelID uint    `json:"hotel_id"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}
