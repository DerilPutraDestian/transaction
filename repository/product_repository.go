package repository

import (
	models "transaction/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll(productCode, search string, limit, offset int) ([]models.Product, int64, error) {
	var products []models.Product
	var count int64

	query := r.db.Model(&models.Product{})

	if productCode != "" {
		query = query.Where("product_code = ?", productCode)
	}

	err := query.Count(&count).
		Preload("Category").
		Limit(limit).
		Offset(offset).
		Find(&products).Error

	return products, count, err
}

func (r *ProductRepository) GetByID(id string) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Category").
		First(&product, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) Create(product *models.Product) error {
	if err := r.db.Create(product).Error; err != nil {
		return err
	}
	return r.db.Preload("Category").First(product, "id = ?", product.ID).Error
}
func (r *ProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(product *models.Product) error {
	return r.db.Delete(product).Error
}
