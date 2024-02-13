package repository

import (
	"gorm.io/gorm"
	"store/models"
)

type ProductPostgres struct {
	db *gorm.DB
}

func NewProductPostgres(db *gorm.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) Create(product models.Product) (uint, error) {
	var newProduct models.Product = models.Product{
		Title:            product.Title,
		Price:            product.Price,
		Description:      product.Description,
		ShortDescription: product.ShortDescription,
		Images:           product.Images,
		Colors:           product.Colors,
		Sizes:            product.Sizes,
		Composition:      product.Composition,
	}
	res := r.db.Create(&newProduct)
	if res.Error != nil {
		return 0, res.Error
	}
	return newProduct.ID, nil
}

func (r *ProductPostgres) Get(page int, limit int) ([]models.Product, int64) {
	var product []models.Product
	var total int64
	r.db.Find(&product).Count(&total)
	r.db.Limit(limit).Offset(page * limit).Find(&product)
	return product, total
}

func (r *ProductPostgres) Delete(ID uint) {
	r.db.Delete(&models.Product{}, ID)
}

func (r *ProductPostgres) Latest() []models.Product {
	var product []models.Product
	r.db.Order("created_at").Limit(5).Find(&product)
	return product
}
