package models

import (
	"gorm.io/gorm"
	"time"
)

type RoomHousekeeping struct {
	gorm.Model
	Description   string    `json:"description"`
	StartDatetime time.Time `json:"start_datetime"`
	Duration      int       `json:"duration"`
	RoomID        uint      `json:"room_id"`
}
