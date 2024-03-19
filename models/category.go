package models

import "gorm.io/gorm"

type CategoryResponse struct {
	Id            int              `json:"category_id"`
	CategoryName  string           `json:"category_name"`
	Subcategories []SubcategoryArr `json:"subcategories" gorm:"foreignKey:CategoryID"`
}

type Category struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
}

type SubcategoryArr struct {
	Id         int
	Name       string
	CategoryID int
}
type Subcategory struct {
	gorm.Model
	Name       string
	CategoryID int
	Category   Category
}

type CategoryInput struct {
	Name string `json:"name" binding:"required"`
}

type SubcategoryInput struct {
	Name       string `json:"name" binding:"required"`
	CategoryID int
	Category   Category
}

type UpdateSubcategoryInput struct {
	Name string `json:"name" binding:"required"`
}
type UpdateCategoryInput struct {
	Name string `json:"name" binding:"required"`
	ID   uint   `json:"id" binding:"required"`
}

type DeleteSubcategoryInput struct {
	ID uint `json:"id" binding:"required"`
}
type DeleteCategoryInput struct {
	ID uint `json:"id" binding:"required"`
}
