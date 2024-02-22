package service

import (
	"store/models"
	"store/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error)
	DeleteUser(ID uint) error
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
	GetUserInfo(ID uint) (models.User, error)
}

type Blog interface {
	Get(page int, limit int) ([]models.Blog, int64)
	Create(blog models.CreateBlogParams) (uint, error)
	Delete(ID uint)
}

type Product interface {
	Get(page int, limit int) ([]models.Product, int64)
	GetById(ID uint) models.Product
	Latest() []models.Product
	Create(product models.Product) (uint, error)
	Delete(ID uint)
	GetCategories() ([]models.Category, error)
	UpdateCategory(ID uint, newName string) error
	DeleteCategory(ID uint) error
	AddCategory(name string) (uint, error)
	GetByName(name string) ([]models.Product, error)
}

type Service struct {
	Authorization
	Product
	Blog
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Blog:          NewBlogService(repos.Blog),
		Product:       NewProductService(repos.Product),
	}
}
