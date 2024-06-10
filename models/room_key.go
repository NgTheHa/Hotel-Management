package models

import (
	"gorm.io/gorm"
	"time"
)

type RoomKey struct {
	gorm.Model
	KeyID    string    `json:"key_id"`
	Barcode  string    `json:"barcode"`
	IssuedAt time.Time `json:"issued_at"`
	IsActive bool      `json:"is_active"`
	IsMaster bool      `json:"is_master"`
	RoomID   uint      `json:"room_id"`
}
