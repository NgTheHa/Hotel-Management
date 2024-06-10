package models

import (
	"gorm.io/gorm"
)

type RoomStyle int

const (
	Standard RoomStyle = iota + 1
	Deluxe
	FamilySuite
	BusinessSuite
)

type RoomStatus int

const (
	Available RoomStatus = iota + 1
	Reserved
	Occupied
	NotAvailable
	BeingServiced
	Other
)

type Room struct {
	gorm.Model
	RoomNumber   string             `json:"room_number"`
	Style        RoomStyle          `json:"style"`
	Status       RoomStatus         `json:"status"`
	BookingPrice float64            `json:"booking_price"`
	IsSmoking    bool               `json:"is_smoking"`
	Keys         []RoomKey          `json:"keys"`
	Housekeeping []RoomHousekeeping `json:"housekeeping"`
}
