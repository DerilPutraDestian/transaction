package service

import (
	models "transaction/model"
	"transaction/repository"
)

type ProductService struct {
	repo        *repository.ProductRepository
	historyRepo repository.HistoryRepository
}

// 🔥 FIX: tambahkan historyRepo di constructor
func NewProductService(repo *repository.ProductRepository, historyRepo repository.HistoryRepository) *ProductService {
	return &ProductService{
		repo:        repo,
		historyRepo: historyRepo,
	}
}

// =========================
// GET
// =========================
func (s *ProductService) ListProducts(id string, search string, limit, offset int) ([]models.Product, int64, error) {
	return s.repo.GetAll(id, search, limit, offset)
}

func (s *ProductService) GetProduct(id string) (*models.Product, error) {
	return s.repo.GetByID(id)
}

// =========================
// CREATE
// =========================
func (s *ProductService) CreateProduct(product *models.Product) error {

	// 🔥 simpan asset dulu
	if err := s.repo.Create(product); err != nil {
		return err
	}

	// 🔥 baru simpan history (SETELAH berhasil)
	_ = s.historyRepo.Create(&models.ProductHistory{
		ProductID:   product.ID,
		Action:      "create",
		Description: "Product created",
	})

	return nil
}

// =========================
// UPDATE
// =========================
func (s *ProductService) UpdateProduct(product *models.Product) error {
	if err := s.repo.Update(product); err != nil {
		return err
	}

	// 🔥 history
	_ = s.historyRepo.Create(&models.ProductHistory{
		ProductID:   product.ID,
		Action:      "update",
		Description: "Product updated",
	})

	return nil
}

// =========================
// DELETE
// =========================
func (s *ProductService) DeleteProduct(product *models.Product) error {

	if err := s.repo.Delete(product); err != nil {
		return err
	}

	// 🔥 history
	_ = s.historyRepo.Create(&models.ProductHistory{
		ProductID:   product.ID,
		Action:      "delete",
		Description: "Product deleted",
	})

	return nil
}
