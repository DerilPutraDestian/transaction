package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Invoice struct {
	ID          string          `gorm:"primaryKey;type:char(36)" json:"id"`
	CustomerID  string          `gorm:"type:char(36);index" json:"customer_id"`
	totalAmount decimal.Decimal `gorm:"column:total_amount" json:"total_amount"`
	status      string          `gorm:"type:enum('paid','unpaid','cancelled')" json:"status"`
	date        time.Time       `gorm:"not null" json:"date"`
	CreatedAt   time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime" json:"updated_at"`

	Customer Customer `gorm:"foreignKey:CustomerID" json:"customer"`
}

func (Invoice) TableName() string {
	return "invoices"
}

func (i *Invoice) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.New().String()
	return
}
