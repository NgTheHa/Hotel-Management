package models

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name string `json:"name"`
	//Address Address `json:"address"`
	Email   string    `json:"email"`
	Phone   string    `json:"phone"`
	Account []Account `json:"account" gorm:"ForeignKey:ID;AssociationForeignKey:ID"`
}

type Guest struct {
	Person
	TotalRoomsCheckedIn int `json:"total_rooms_checked_in"`
}

type Receptionist struct {
	Person
}

func (r *Receptionist) CreateBooking() {
	// Logic to create booking
}

type Server struct {
	Person
}
