package repository

import (
	"gorm.io/gorm"
	"store/models"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error)
	GetUser(email, password string) (models.User, error)
	GetUserInfo(ID uint) (models.User, error)
}

type Blog interface {
	Get(page int, limit int) ([]models.Blog, int64)
	Create(blog models.CreateBlogParams) (uint, error)
	Delete(ID uint)
}

type Product interface {
	Get(page int, limit int) ([]models.Product, int64)
	GetById(id uint) models.Product
	Latest() []models.Product
	Create(product models.Product) (uint, error)
	Delete(ID uint)
	GetCategories() ([]models.Category, error)
	UpdateCategory(ID uint, newName string) error
	DeleteCategory(ID uint) error
	AddCategory(name string) (uint, error)
	GetByName(name string) ([]models.Product, error)
}

type Repository struct {
	Authorization
	Product
	Blog
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Blog: NewBlogPostgres(db), Product: NewProductPostgres(db), Authorization: NewAuthPostgres(db)}
}
