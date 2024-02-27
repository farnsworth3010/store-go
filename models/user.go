package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email       string `json:"email" binding:"required" gorm:"uniqueIndex:email,where:deleted_at is null"`
	Firstname   string `json:"firstname" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Password    string `json:"password" binding:"required"`
	RoleID      int    `json:"role_id" gorm:"default:2"`
	Role        Role
}
type ShortUser struct {
	gorm.Model
	Email       string `json:"email" binding:"required" gorm:"uniqueIndex:email,where:deleted_at is null"`
	Firstname   string `json:"firstname" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	RoleID      int    `json:"role_id" gorm:"default:2"`
	Role        Role
}

type Role struct {
	gorm.Model
	Name string `json:"name"`
}
