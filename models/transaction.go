package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Amount          float64   `json:"amount"`
	Description     string    `json:"description"`
	Category        string    `json:"category"`
	Date            time.Time `json:"date"`
	TransactionType string    `json:"transaction_type"`
	UserID          uint      `json:"user_id"`
}
