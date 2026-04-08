package service

import (
	models "transaction/model"
	"transaction/repository"
)

type CustomerService struct {
	repo *repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) ListCustomers(search string, limit, offset int) ([]models.Customer, int64, error) {
	return s.repo.GetAll(search, limit, offset)
}

func (s *CustomerService) GetCustomer(id uint) (*models.Customer, error) {
	return s.repo.GetByID(id)
}

func (s *CustomerService) CreateCustomer(c *models.Customer) error {
	return s.repo.Create(c)
}

func (s *CustomerService) UpdateCustomer(c *models.Customer) error {
	return s.repo.Update(c)
}

func (s *CustomerService) DeleteCustomer(id uint) error {
	return s.repo.Delete(id)
}
