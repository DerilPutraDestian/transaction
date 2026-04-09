package service

import (
	models "transaction/model"
	"transaction/repository"
)

type InvoiceItemService struct {
	repo *repository.InvoiceItemRepository
}

func NewInvoiceItemService(repo *repository.InvoiceItemRepository) *InvoiceItemService {
	return &InvoiceItemService{repo: repo}
}
func (s *InvoiceItemService) ListInvoiceItems(invoiceItemID string, search string, limit, offset int) ([]models.InvoiceItem, int64, error) {
	return s.repo.GetAll(invoiceItemID, search, limit, offset)
}
func (s *InvoiceItemService) GetInvoiceItemByID(id string) (*models.InvoiceItem, error) {
	return s.repo.GetByID(id)
}
func (s *InvoiceItemService) CreateInvoiceItem(invoiceItem *models.InvoiceItem) error {
	return s.repo.Create(invoiceItem)
}

func (s *InvoiceItemService) UpdateInvoiceItem(invoiceItem *models.InvoiceItem) error {
	return s.repo.Update(invoiceItem)
}
