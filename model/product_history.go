package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductHistory struct {
	ID          string    `gorm:"primaryKey;type:char(36)" json:"id"`
	ProductID   string    `gorm:"type:char(36);index" json:"product_id"`
	Action      string    `gorm:"size:50" json:"action"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (h *ProductHistory) BeforeCreate(tx *gorm.DB) (err error) {
	h.ID = uuid.New().String()
	return
}
