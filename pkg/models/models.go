package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `json:"name"`
	PhoneNumber string `json:"phone number" gorm:"unique"`
	Email       string `json:"email" gorm:"unique"`
	Password    []byte `json:"password"`
}
