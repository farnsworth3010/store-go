package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
}

type CategoryInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategoryInput struct {
	Name string `json:"name" binding:"required"`
	ID   uint   `json:"id" binding:"required"`
}

type DeleteCategoryInput struct {
	ID uint `json:"id" binding:"required"`
}
