package repository

import (
	models "transaction/model"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) GetAll(assetLoanCode, search string, limit, offset int) ([]models.Transaction, int64, error) {
	var assetsLoan []models.Transaction
	var count int64

	query := r.db.Model(&models.Transaction{})

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
		Find(&assetsLoan).Error

	return assetsLoan, count, err
}

func (r *TransactionRepository) GetByID(id string) (*models.Transaction, error) {
	var data models.Transaction
	// Di sini juga sebaiknya tambahkan Preload("Product.Category") agar data lengkap saat ambil detail
	err := r.db.Preload("Product").Preload("Product.Category").Preload("Customer").First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *TransactionRepository) Create(transaction *models.Transaction) error {
	return r.db.Omit("Product", "Customer").Create(transaction).Error
}

func (r *TransactionRepository) Update(transaction *models.Transaction) error {
	return r.db.Save(transaction).Error
}
