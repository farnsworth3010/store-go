package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email       string `json:"email" binding:"required"`
	Firstname   string `json:"firstname" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
