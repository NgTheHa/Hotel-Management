package models

import (
	"gorm.io/gorm"
	"time"
)

type Notification struct {
	gorm.Model
	NotificationID uint      `json:"notification_id" gorm:"primaryKey"`
	CreatedOn      time.Time `json:"created_on"`
	Content        string    `json:"content"`
	Send           bool      `json:"send"`
	RoomBookingID  uint      `json:"room_booking_id"`
}
