package repository

import (
	"gorm.io/gorm"
	"store/models"
)

type BlogPostgres struct {
	db *gorm.DB
}

func NewBlogPostgres(db *gorm.DB) *BlogPostgres {
	return &BlogPostgres{db: db}
}

func (r *BlogPostgres) Create(blog models.CreateBlogParams) (uint, error) {
	var newBlog models.Blog = models.Blog{Title: blog.Title, Text: blog.Text}
	res := r.db.Create(&newBlog)
	if res.Error != nil {
		return 0, res.Error
	}
	return newBlog.ID, nil
}

func (r *BlogPostgres) Get(page int, limit int) ([]models.Blog, int64) {
	var blog []models.Blog
	var total int64
	r.db.Find(&blog).Count(&total)
	r.db.Limit(limit).Offset(page * limit).Find(&blog)
	return blog, total
}

func (r *BlogPostgres) Delete(ID uint) {
	r.db.Delete(&models.Blog{}, ID)
}
