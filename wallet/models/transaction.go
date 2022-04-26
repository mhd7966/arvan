package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model

	UserID      int    `json:"user_id"`
	Code        string `json:"code"`
	CodeType    int    `json:"code_type"`
	Value       int    `json:"value"`
	ValueType   int    `json:"value_type"`
	InitBalance int    `json:"init_balance"`
	NewBalance  int    `json:"new_balance"`
}
