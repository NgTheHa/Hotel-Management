package models

import (
	"time"

	"gorm.io/gorm"
)

type BookingStatus int

const (
	Requested BookingStatus = iota + 1
	Confirmed
	CheckedIn
	CheckedOut
)

type RoomBooking struct {
	gorm.Model
	ReservationNumber string        `json:"reservation_number" gorm:"uniqueIndex;not null"`
	StartDate         time.Time     `gorm:"not null;index"`
	DurationInDays    int           `gorm:"not null;index"`
	BookingStatus     BookingStatus `gorm:"not null"`
	CheckIn           *time.Time
	CheckOut          *time.Time
	GuestID           uint
	RoomID            uint
	Room              Room           `gorm:"foreignKey:RoomID"`
	Payments          []Payment      `gorm:"foreignKey:BookingID"`
	Notifications     []Notification `gorm:"foreignKey:RoomBookingID"`
}

// NewRoomBooking tạo một đặt phòng mới
func NewRoomBooking(reservationNumber string, startDate time.Time, durationInDays int, bookingStatus BookingStatus, guestID uint, roomID uint) *RoomBooking {
	return &RoomBooking{
		ReservationNumber: reservationNumber,
		StartDate:         startDate,
		DurationInDays:    durationInDays,
		BookingStatus:     bookingStatus,
		GuestID:           guestID,
		RoomID:            roomID,
	}
}

// UpdateBookingStatus cập nhật trạng thái đặt phòng
func (r *RoomBooking) UpdateBookingStatus(status BookingStatus) {
	r.BookingStatus = status
}

// SetCheckIn thiết lập thời gian check-in
func (r *RoomBooking) SetCheckIn(checkIn time.Time) {
	r.CheckIn = &checkIn
}

// SetCheckOut thiết lập thời gian check-out
func (r *RoomBooking) SetCheckOut(checkOut time.Time) {
	r.CheckOut = &checkOut
}

// CalculateTotalPrice tính toán tổng giá trị đặt phòng
func (r *RoomBooking) CalculateTotalPrice() float64 {
	// Logic tính toán tổng giá trị dựa trên số ngày và giá của phòng
	return 0
}
