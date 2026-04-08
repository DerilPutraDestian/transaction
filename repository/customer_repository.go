package repository

import (
	models "transaction/model"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) GetAll(search string, limit, offset int) ([]models.Customer, int64, error) {
	var customers []models.Customer
	var total int64

	query := r.db.Model(&models.Customer{})

	if search != "" {
		// Menggunakan kolom customer_name sesuai database kamu
		query = query.Where("customer_name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	err := query.Count(&total).Limit(limit).Offset(offset).Find(&customers).Error
	return customers, total, err
}

func (r *CustomerRepository) GetByID(id uint) (*models.Customer, error) {
	var customer models.Customer
	// Mencari berdasarkan primary key (customer_id)
	err := r.db.First(&customer, id).Error
	return &customer, err
}

func (r *CustomerRepository) Create(c *models.Customer) error {
	return r.db.Create(c).Error
}

func (r *CustomerRepository) Update(c *models.Customer) error {
	return r.db.Save(c).Error
}

func (r *CustomerRepository) Delete(id uint) error {
	return r.db.Delete(&models.Customer{}, id).Error
}
