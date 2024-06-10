package models

import (
	"gorm.io/gorm"
	"time"
)

type Notification struct {
	gorm.Model
	NotificationID int
	CreatedOn      time.Time `json:"created_on"`
	Content        string    `json:"content"`
	Sent           bool      `json:"sent"`
}
