package models

import "gorm.io/gorm"

type PaymentStatus string

const (
	Unpaid    PaymentStatus = "UNPAID"
	Pending   PaymentStatus = "PENDING"
	Completed PaymentStatus = "COMPLETED"
	Failed    PaymentStatus = "FAILED"
	Declined  PaymentStatus = "DECLINED"
	Cancelled PaymentStatus = "CANCELLED"
	Abandoned PaymentStatus = "ABANDONED"
	Settling  PaymentStatus = "SETTLING"
	Settled   PaymentStatus = "SETTLED"
	Refunded  PaymentStatus = "REFUNDED"
)

type Payment struct {
	gorm.Model
	Amount    float64       `json:"amount" gorm:"not null"`
	Status    PaymentStatus `json:"status" gorm:"not null"`
	Method    string        `json:"method" gorm:"not null"`
	BookingID uint
}
