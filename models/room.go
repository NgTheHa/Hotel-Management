package models

import (
	"gorm.io/gorm"
)

type RoomStyle string

const (
	Standard      RoomStyle = "STANDARD"
	Deluxe        RoomStyle = "DELUXE"
	FamilySuite   RoomStyle = "FAMILY_SUITE"
	BusinessSuite RoomStyle = "BUSINESS_SUITE"
)

type RoomStatus string

const (
	Available     RoomStatus = "AVAILABLE"
	Reserved      RoomStatus = "RESERVED"
	Occupied      RoomStatus = "OCCUPIED"
	NotAvailable  RoomStatus = "NOT_AVAILABLE"
	BeingServiced RoomStatus = "BEING_SERVICED"
)

type Room struct {
	gorm.Model
	RoomNumber string  `json:"room_number"`
	Style      string  `json:"style"`  // Mô hình đơn giản hóa không sử dụng enum
	Status     string  `json:"status"` // Mô hình đơn giản hóa không sử dụng enum
	Price      float64 `json:"price"`
	IsSmoking  bool    `json:"is_smoking"`
	LocationID uint    `json:"location_id" gorm:"foreignKey:ID"` // Foreign key
}
