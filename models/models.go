package models

import "gorm.io/gorm"

// User has one CreditCard, UserID is the foreign key
type User struct {
	gorm.Model
	Name         string     `json:"name"`
	CreditCard   CreditCard `json:"creditCard"`
	CreditCardId uint       `gorm:"foreignKey:ID" json:"-"`
}

type CreditCard struct {
	gorm.Model
	Number string `json:"number"`
}
