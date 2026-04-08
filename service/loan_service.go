package service

import (
	models "transaction/model"
	"transaction/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}
func (s *TransactionService) ListTransactions(transactionID string, search string, limit, offset int) ([]models.Transaction, int64, error) {
	return s.repo.GetAll(transactionID, search, limit, offset)
}
func (s *TransactionService) GetTransactionByID(id string) (*models.Transaction, error) {
	return s.repo.GetByID(id)
}
func (s *TransactionService) CreateTransaction(t *models.Transaction) error {
	return s.repo.Create(t)
}

func (s *TransactionService) UpdateTransaction(t *models.Transaction) error {
	return s.repo.Update(t)
}
