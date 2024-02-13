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
