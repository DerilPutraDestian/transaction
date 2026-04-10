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

// GetAll sekarang menggunakan customerID dan status sebagai filter
func (r *InvoiceRepository) GetAll(customerID, status string, limit, offset int) ([]models.Invoice, int64, error) {
	var invoices []models.Invoice
	var count int64

	query := r.db.Model(&models.Invoice{})

	// Filter berdasarkan CustomerID jika disediakan
	if customerID != "" {
		query = query.Where("customer_id = ?", customerID)
	}

	// Filter berdasarkan Status jika disediakan
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Hitung total record untuk pagination
	query.Count(&count)

	err := query.
		// Pilih kolom yang ada di struct Invoice (Huruf besar)
		Select("id", "customer_id", "total_amount", "status", "date", "created_at", "updated_at").
		// Preload Customer secara spesifik agar ringan
		Preload("Customer", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name") // Pastikan field 'name' ada di tabel customers
		}).
		Limit(limit).
		Offset(offset).
		Find(&invoices).Error

	return invoices, count, err
}

func (r *InvoiceRepository) GetByID(id string) (*models.Invoice, error) {
	var data models.Invoice
	// Preload Customer lengkap untuk detail
	err := r.db.Preload("Customer").First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *InvoiceRepository) Create(invoice *models.Invoice) error {
	// Omit "Customer" agar tidak mencoba insert/update ke tabel customers
	return r.db.Omit("Customer").Create(invoice).Error
}

func (r *InvoiceRepository) Update(invoice *models.Invoice) error {
	// Omit "Customer" agar data referensi tidak ikut terupdate secara tidak sengaja
	return r.db.Omit("Customer").Save(invoice).Error
}
