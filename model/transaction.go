package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID         string `gorm:"primaryKey;type:char(36)" json:"id"`
	ProductID  string `gorm:"type:char(36);index;not null" json:"product_id"`
	CustomerID string `gorm:"type:char(36);index;not null" json:"customer_id"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Product  Product  `gorm:"foreignKey:ProductID;->:false" json:"product"`
	Customer Customer `gorm:"foreignKey:CustomerID;->:false" json:"customer"`
}

func (Transaction) TableName() string {
	return "transactions"
}

func (l *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	l.ID = uuid.New().String()
	return
}
