package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title string `json:"title" binding:"required"`
	Text  string `json:"text" binding:"required"`
}

type ShortBlog struct {
	gorm.Model
	Title string `json:"title" binding:"required"`
}

type CreateBlogParams struct {
	Title string `json:"title" binding:"required"`
	Text  string `json:"text" binding:"required"`
}
