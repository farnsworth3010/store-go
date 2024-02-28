package repository

import (
	"store/models"

	"gorm.io/gorm"
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
		BrandID:          product.BrandID,
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
func (r *ProductPostgres) GetById(id uint) models.Product {
	var product models.Product
	r.db.Find(&product, id)
	return product
}

func (r *ProductPostgres) Delete(ID uint) {
	r.db.Delete(&models.Product{}, ID)
}

func (r *ProductPostgres) Latest() []models.Product {
	var product []models.Product
	r.db.Order("created_at").Limit(5).Find(&product)
	return product
}

func (r *ProductPostgres) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	res := r.db.Find(&categories)
	if res.Error != nil {
		return categories, res.Error
	}
	return categories, nil
}

func (r *ProductPostgres) UpdateCategory(ID uint, newName string) error {
	res := r.db.Model(&models.Category{
		Model: gorm.Model{ID: ID},
	}).Update("name", newName)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *ProductPostgres) DeleteCategory(ID uint) error {
	res := r.db.Delete(&models.Category{}, ID)
	if res.Error != nil {
		return res.Error
	}
	return res.Error
}

func (r *ProductPostgres) AddCategory(name string) (uint, error) {
	var newCategory = models.Category{
		Name: name,
	}
	res := r.db.Create(&newCategory)
	if res.Error != nil {
		return 0, res.Error
	}
	return newCategory.ID, nil
}

func (r *ProductPostgres) GetByName(name string) ([]models.Product, error) {
	var product []models.Product
	res := r.db.Where("title LIKE ?", "%"+name+"%").Find(&product)
	if res.Error != nil {
		return product, res.Error
	}
	return product, nil
}
func (r *ProductPostgres) GetBrands() ([]models.Brand, error) {
	var brands []models.Brand
	res := r.db.Find(&brands)
	if res.Error != nil {
		return brands, res.Error
	}
	return brands, nil
}
