package models

import (
	"time"

	"gorm.io/gorm"
)

type ChargeCode struct {
	gorm.Model
	Name           string    `json:"name"`
	Value          int       `json:"value"`
	Capacity       int       `json:"capacity"`
	MaxCapacity    int       `json:"max_capacity"`
	ExpirationDate time.Time `json:"expiration_date"`
}
