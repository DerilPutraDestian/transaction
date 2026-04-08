package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID         string          `gorm:"primaryKey;type:char(36)" json:"id"`
	Name       string          `gorm:"size:255;not null" json:"name"`
	price      decimal.Decimal `gorm:"column:price" json:"price"`
	CategoryID string          `gorm:"type:char(36);index" json:"category_id"`
	CreatedAt  time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time       `gorm:"autoUpdateTime" json:"updated_at"`

	Category  Category         `gorm:"foreignKey:CategoryID" json:"category"`
	Histories []ProductHistory `gorm:"foreignKey:ProductID" json:"histories,omitempty"`
}

func (Product) TableName() string {
	return "products"
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return
}
