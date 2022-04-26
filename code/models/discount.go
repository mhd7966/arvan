package models

import (
	"gorm.io/gorm"
)

type Discount struct {
	gorm.Model
	Name           string `json:"name"`
	Capacity       int    `json:"capacity"`
	MaxCapacity    int    `json:"max_capacity"`
	ExpirationDate int    `json:"expiration_date"`
	MaxDiscount    int    `json:"max_discount"`
	MinPurchase    int    `json:"min_purchase"`
	Percent        int    `json:"percent"`
}
