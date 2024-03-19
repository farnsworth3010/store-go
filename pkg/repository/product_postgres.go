package repository

import (
	"database/sql"
	"encoding/json"
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

func (r *ProductPostgres) GetCategories() ([]models.CategoryResponse, error) {
	var categories []models.CategoryResponse

	rows, err := r.db.Raw("SELECT categories.id as category_id, categories.name AS category_name, json_agg(json_build_object('id', subcategories.id, 'name', subcategories.name)) as subcategories FROM categories LEFT JOIN subcategories ON categories.id=subcategories.category_id GROUP BY categories.id, categories.name").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category models.CategoryResponse
		var subcategoriesJSON sql.RawBytes
		if err := rows.Scan(&category.Id, &category.CategoryName, &subcategoriesJSON); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(subcategoriesJSON, &category.Subcategories); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, err
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

func (r *ProductPostgres) Update(product models.Product) error {
	res := r.db.Model(&models.Product{}).Where("id = ?", product.ID).Updates(&models.Product{Title: product.Title, Price: product.Price, Description: product.Description, ShortDescription: product.ShortDescription, Images: product.Images, Colors: product.Colors, Sizes: product.Sizes, Composition: product.Composition, BrandID: product.BrandID})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func BrandFilter(brand_id int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if brand_id != 0 {
			return db.Where("brand_id = ?", brand_id)
		}
		return db
	}
}

func TitleFilter(title string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if title != "" {
			return db.Where("title LIKE ?", "%"+title+"%")
		}
		return db
	}
}

func SortFilter(criterion int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch criterion {
		case 1:
			return db
		case 2:
			return db.Order("price ASC")
		case 3:
			return db.Order("price DESC")
		case 4:
			return db.Order("created_at DESC")
		default:
			return db
		}
	}
}

func (r *ProductPostgres) Filter(filters models.ProductFilters, page int, limit int) ([]models.Product, int64) {
	var products []models.Product
	var total int64

	r.db.Scopes(TitleFilter(filters.Title), BrandFilter(filters.BrandID), SortFilter(filters.SortCriterion)).Find(&products).Limit(limit).Offset(page * limit)

	r.db.Scopes(TitleFilter(filters.Title), BrandFilter(filters.BrandID),
		SortFilter(filters.SortCriterion)).Find(&models.Product{}).Count(&total)

	return products, total
}
