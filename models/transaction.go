package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model

	TransactionType string  `json:"transactionType" gorm:"type:varchar(20);not null"`
	Amount          float64 `json:"amount" gorm:"type:decimal(10,2);not null"`
	Description     string  `json:"description" gorm:"type:varchar(255);not null"`
	Category        string  `json:"category" gorm:"type:varchar(50);not null"`
	Date            string  `json:"date" gorm:"type:varchar(20);not null"`
	UserID          uint    `json:"user_id"`
}
