package repository

import (
	"store/models"

	"gorm.io/gorm"
)

type BrandPostgres struct {
	db *gorm.DB
}

func NewBrandPostgres(db *gorm.DB) *BrandPostgres {
	return &BrandPostgres{db: db}
}

func (r *BrandPostgres) Create(brand models.CreateBrandParams) (uint, error) {
	var newBrand models.Brand = models.Brand{Name: brand.Name}
	res := r.db.Create(&newBrand)
	if res.Error != nil {
		return 0, res.Error
	}
	return newBrand.ID, nil
}

func (r *BrandPostgres) Update(brand models.EditBrandParams) error {
	var newBrand models.Brand = models.Brand{Name: brand.Name}
	res := r.db.Model(&models.Brand{}).Where("id = ?", brand.ID).Updates(&newBrand)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *BrandPostgres) Get() ([]models.Brand, int64) {
	var brand []models.Brand
	var total int64
	r.db.Find(&brand).Count(&total)
	r.db.Find(&brand)
	return brand, total
}

func (r *BrandPostgres) Delete(ID uint) {
	r.db.Delete(&models.Brand{}, ID)
}
