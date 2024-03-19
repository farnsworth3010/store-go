package service

import (
	"store/models"
	"store/pkg/repository"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(blog models.Product) (uint, error) {
	return s.repo.Create(blog)
}

func (s *ProductService) Get(page int, limit int) ([]models.Product, int64) {
	return s.repo.Get(page, limit)
}

func (s *ProductService) Delete(ID uint) {
	s.repo.Delete(ID)
}
func (s *ProductService) Latest() []models.Product {
	return s.repo.Latest()
}

func (s *ProductService) GetById(id uint) models.Product {
	return s.repo.GetById(id)
}

func (s *ProductService) GetCategories() ([]models.CategoryResponse, error) {
	return s.repo.GetCategories()
}
func (s *ProductService) UpdateCategory(ID uint, newName string) error {
	return s.repo.UpdateCategory(ID, newName)
}
func (s *ProductService) DeleteCategory(ID uint) error {
	return s.repo.DeleteCategory(ID)
}
func (s *ProductService) AddCategory(name string) (uint, error) {
	return s.repo.AddCategory(name)
}
func (s *ProductService) GetByName(name string) ([]models.Product, error) {
	return s.repo.GetByName(name)
}

func (s *ProductService) GetBrands() ([]models.Brand, error) {
	return s.repo.GetBrands()
}

func (s *ProductService) Update(product models.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) Filter(filters models.ProductFilters, page int, limit int) ([]models.Product, int64) {
	return s.repo.Filter(filters, page, limit)
}
