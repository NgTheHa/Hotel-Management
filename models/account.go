package models

import (
	"gorm.io/gorm"
)

type AccountStatus string

const (
	Active      AccountStatus = "ACTIVE"
	Closed      AccountStatus = "CLOSED"
	Canceled    AccountStatus = "CANCELED"
	Blacklisted AccountStatus = "BLACKLISTED"
)

type Account struct {
	gorm.Model
	ID       uint          `json:"id" gorm:"primaryKey"`
	Password string        `json:"password"`
	Status   AccountStatus `json:"status"`
}

//func (a *Account) ResetPassword() {
//	// Logic to reset password
//	//	TODO
//}
