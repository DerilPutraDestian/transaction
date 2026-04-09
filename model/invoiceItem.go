package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type InvoiceItem struct {
	ID        string          `gorm:"primaryKey;type:char(36)" json:"id"`
	ProductID string          `gorm:"type:char(36);index" json:"product_id"`
	UnitPrice decimal.Decimal `gorm:"column:unit_price" json:"unit_price"`
	Subtotal  decimal.Decimal `gorm:"column:subtotal" json:"subtotal"` // Huruf kapital agar exportable
	Qty       int             `gorm:"not null" json:"qty"`
	InvoiceID string          `gorm:"type:char(36);index" json:"invoice_id"`
	CreatedAt time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime" json:"updated_at"`

	// Relasi
	Product Product `gorm:"foreignKey:ProductID" json:"product"`
	Invoice Invoice `gorm:"foreignKey:InvoiceID" json:"invoice"`
}

func (InvoiceItem) TableName() string {
	return "invoice_items"
}

func (i *InvoiceItem) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.New().String()
	// Hitung subtotal otomatis sebelum simpan
	i.Subtotal = i.UnitPrice.Mul(decimal.NewFromInt(int64(i.Qty)))
	return
}
