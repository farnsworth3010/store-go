package service

import (
	"store/models"
	"store/pkg/repository"
)

type BrandService struct {
	repo repository.Brand
}

func NewBrandService(repo repository.Brand) *BrandService {
	return &BrandService{repo: repo}
}

func (s *BrandService) Create(brand models.CreateBrandParams) (uint, error) {
	return s.repo.Create(brand)
}
func (s *BrandService) Update(brand models.EditBrandParams) error {
	return s.repo.Update(brand)
}

func (s *BrandService) Get() ([]models.Brand, int64) {
	return s.repo.Get()
}

func (s *BrandService) Delete(ID uint) {
	s.repo.Delete(ID)
}
