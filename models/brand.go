package models

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	name string
}
