package repository

import (
	"store/models"

	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error)
	GetUser(email, password string) (models.User, error)
	GetUserInfo(ID uint) (models.User, error)
	DeleteUser(ID uint) error
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

type Panel interface {
	GetAdmins() ([]models.Admin, error)
	GetUsers() ([]models.ShortUser, error)
	GetBlogs() ([]models.ShortBlog, error)
	SetRole(ID uint, RoleID uint) error
}

type Repository struct {
	Authorization
	Blog
	Panel
	Product
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Blog: NewBlogPostgres(db), Product: NewProductPostgres(db), Authorization: NewAuthPostgres(db), Panel: NewPanelPostgres(db)}
}
