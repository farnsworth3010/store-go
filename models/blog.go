package models

type Blog struct {
	Id            int    `json:"-" db:"id"`
	Title         string `json:"title" binding:"required"`
	Text          string `json:"text" binding:"required"`
	Creation_Date string `json:"creation_date"`
}

type CreateBlogParams struct {
	Title string `json:"title" binding:"required"`
	Text  string `json:"text" binding:"required"`
}
