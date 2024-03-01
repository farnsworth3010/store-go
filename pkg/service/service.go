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
	Update(product models.Product) error
	GetCategories() ([]models.Category, error)
	UpdateCategory(ID uint, newName string) error
	DeleteCategory(ID uint) error
	AddCategory(name string) (uint, error)
	GetByName(name string) ([]models.Product, error)
	GetBrands() ([]models.Brand, error)
}

type Panel interface {
	GetAdmins() ([]models.Admin, error)
	GetUsers() ([]models.ShortUser, error)
	GetBlogs() ([]models.ShortBlog, error)
	SetRole(ID uint, RoleID uint) error
}

type Service struct {
	Authorization
	Blog
	Panel
	Product
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Blog:          NewBlogService(repos.Blog),
		Product:       NewProductService(repos.Product),
		Panel:         NewPanelService(repos.Panel),
	}
}
