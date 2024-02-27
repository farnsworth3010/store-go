package models

type Admin struct {
	ID          int    `json:"ID"`
	Email       string `json:"email" binding:"required" gorm:"uniqueIndex:email,where:deleted_at is null"`
	Firstname   string `json:"firstname" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	RoleID      int    `json:"role_id" gorm:"default:2"`
	Role        Role
}

type SetRoleInput struct {
	ID     int `json:"ID"`
	RoleID int `json:"role_id"`
}
