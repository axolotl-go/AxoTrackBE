package models

import (
	"gorm.io/gorm"
)

type SignUp struct {
	gorm.Model

	Username string `json:"username" gorm:"type:varchar(50);uniqueIndex;not null"`
	Email    string `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	Phone    string `json:"phone" gorm:"type:varchar(20);uniqueIndex;not null"`
	Image    string `json:"image" gorm:"type:varchar(255);not null"`
}

type SignIn struct {
	gorm.Model

	Email    string `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
}

type UserID struct {
	gorm.Model

	Username    string        `json:"username" gorm:"type:varchar(50);uniqueIndex;not null"`
	Email       string        `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Image       string        `json:"image" gorm:"type:varchar(255);not null"`
	Transaction []Transaction `json:"transaction" `
}
