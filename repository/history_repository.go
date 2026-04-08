package repository

import (
	models "transaction/model"

	"gorm.io/gorm"
)

type HistoryRepository interface {
	Create(h *models.ProductHistory) error
}

type historyRepo struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	return &historyRepo{db}
}

func (r *historyRepo) Create(h *models.ProductHistory) error {
	return r.db.Create(h).Error
}
