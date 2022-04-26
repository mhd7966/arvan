package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PhoneNumber string `json:"phone_number"`
	Balance     int    `json:"balance"`
}
