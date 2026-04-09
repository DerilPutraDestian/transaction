package service

import (
	models "transaction/model"
	"transaction/repository"
)

type InvoiceService struct {
	repo *repository.InvoiceRepository
}

func NewInvoiceService(repo *repository.InvoiceRepository) *InvoiceService {
	return &InvoiceService{repo: repo}
}
func (s *InvoiceService) ListInvoices(invoiceID string, search string, limit, offset int) ([]models.Invoice, int64, error) {
	return s.repo.GetAll(invoiceID, search, limit, offset)
}
func (s *InvoiceService) GetInvoiceByID(id string) (*models.Invoice, error) {
	return s.repo.GetByID(id)
}
func (s *InvoiceService) CreateInvoice(invoice *models.Invoice) error {
	return s.repo.Create(invoice)
}

func (s *InvoiceService) UpdateInvoice(invoice *models.Invoice) error {
	return s.repo.Update(invoice)
}
