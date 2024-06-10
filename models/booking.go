package models

import (
	"gorm.io/gorm"
	"time"
)

type Booking struct {
	gorm.Model
	RoomID       uint      `json:"room_id"`
	CustomerName string    `json:"customer_name"`
	CheckIn      time.Time `json:"check_in"`
	CheckOut     time.Time `json:"check_out"`
	Status       string    `json:"status"`
	PaymentID    uint      `json:"payment_id"`
	Payments     []Payment `json:"payments"`
}
