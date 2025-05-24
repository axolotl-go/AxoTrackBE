package models

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model

	ClientName     string    `json:"borrower_name"`
	Amount         float64   `json:"amount"`
	InterestRate   float64   `json:"interest_rate"`
	DurationMonths int       `json:"duration_months"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	Status         string    `json:"status"`
	UserID         uint      `json:"user_id"`
}

type LoanDeposit struct {
	gorm.Model

	ClientID uint    `json:"user_id"`
	Amount   float64 `json:"amount"`
}
