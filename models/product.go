package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title            string         `json:"title" binding:"required"`
	Price            int            `json:"price" binding:"required"`
	Description      string         `json:"description" binding:"required"`
	ShortDescription string         `json:"short_description" binding:"required"`
	Images           pq.StringArray `json:"images" gorm:"type:text[]"`
	Colors           pq.StringArray `json:"colors" gorm:"type:varchar(20)[]"`
	Sizes            pq.StringArray `json:"sizes" gorm:"type:varchar(20)[]"`
	Composition      pq.StringArray `json:"composition" gorm:"type:varchar(20)[]"`
	Subcategories    []Subcategory  `gorm:"many2many:product_subcategories;"`
	BrandID          int            `json:"brand_id" binding:"required"`
	Brand            Brand
}

type SearchProductInput struct {
	Name string `json:"name" binding:"required"`
}
