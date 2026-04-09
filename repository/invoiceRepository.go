package repository

import (
	models "transaction/model"

	"gorm.io/gorm"
)

type InvoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

func (r *InvoiceRepository) GetAll(assetLoanCode, search string, limit, offset int) ([]models.Invoice, int64, error) {
	var invoices []models.Invoice
	var count int64

	query := r.db.Model(&models.Invoice{})

	if assetLoanCode != "" {
		query = query.Where("asset_loan_code = ?", assetLoanCode)
	}

	// Gunakan search jika ada (opsional, tergantung kebutuhanmu)
	if search != "" {
		// Contoh: mencari berdasarkan nama asset melalui join atau search langsung jika ada fieldnya
	}

	err := query.Count(&count).
		Preload("Product").          // Load data Product
		Preload("Product.Category"). // Load data Category MILIK Product (Nested Preload)
		Preload("Customer").         // Load data Customer yang meminjam
		Limit(limit).
		Offset(offset).
		Find(&invoices).Error

	return invoices, count, err
}

func (r *InvoiceRepository) GetByID(id string) (*models.Invoice, error) {
	var data models.Invoice
	// Di sini juga sebaiknya tambahkan Preload("Product.Category") agar data lengkap saat ambil detail
	err := r.db.Preload("Product").Preload("Product.Category").Preload("Customer").First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *InvoiceRepository) Create(invoice *models.Invoice) error {
	return r.db.Omit("Product", "Customer").Create(invoice).Error
}

func (r *InvoiceRepository) Update(invoice *models.Invoice) error {
	return r.db.Save(invoice).Error
}
