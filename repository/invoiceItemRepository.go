package repository

import (
	"errors"
	models "transaction/model"

	"gorm.io/gorm"
)

type InvoiceItemRepository struct {
	db *gorm.DB
}

func NewInvoiceItemRepository(db *gorm.DB) *InvoiceItemRepository {
	return &InvoiceItemRepository{db: db}
}

func (r *InvoiceItemRepository) GetAll(id, search string, limit, offset int) ([]models.InvoiceItem, int64, error) {
	var invoiceItems []models.InvoiceItem
	var count int64

	query := r.db.Model(&models.InvoiceItem{})

	if id != "" {
		query = query.Where("id = ?", id)
	}

	err := query.Count(&count).
		Preload("Product.Category"). // Jalur: InvoiceItem -> Product -> Category
		Preload("Invoice.Customer"). // Jalur: InvoiceItem -> Invoice -> Customer
		Limit(limit).
		Offset(offset).
		Find(&invoiceItems).Error

	return invoiceItems, count, err
}

func (r *InvoiceItemRepository) GetByID(id string) (*models.InvoiceItem, error) {
	var data models.InvoiceItem
	err := r.db.
		Preload("Product.Category").
		Preload("Invoice.Customer").
		First(&data, "id = ?", id).Error
	return &data, err
}

func (r *InvoiceItemRepository) Create(item *models.InvoiceItem) error {
	// Memulai Transaksi Database
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 1. Cek & Ambil Produk (Lock record untuk keamanan concurrency jika perlu)
		var product models.Product
		if err := tx.First(&product, "id = ?", item.ProductID).Error; err != nil {
			return errors.New("produk tidak ditemukan")
		}

		// 2. Validasi Stok
		if product.Stock < item.Qty {
			return errors.New("stok produk tidak mencukupi")
		}
		newStock := product.Stock - item.Qty
		if err := tx.Model(&product).Update("stock", newStock).Error; err != nil {
			return errors.New("gagal memperbarui stok produk")
		}
		if err := tx.Omit("Product", "Invoice").Create(item).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *InvoiceItemRepository) Update(invoiceItem *models.InvoiceItem) error {
	return r.db.Save(invoiceItem).Error
}
