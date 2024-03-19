package models

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	Name string
}

type CreateBrandParams struct {
	Name string `json:"name" binding:"required"`
}

type EditBrandParams struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
