package handlers

import (
	"strconv"
	models "transaction/model"
	"transaction/service"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(s *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: s}
}

// Index mengambil daftar transaksi dengan paginasi
func (h *TransactionHandler) Index(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	data, total, err := h.service.ListTransactions("", "", limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal mengambil data transaksi",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"total":  total,
		"data":   data,
	})
}

// Store membuat transaksi baru
func (h *TransactionHandler) Store(c *fiber.Ctx) error {
	req := new(models.Transaction) // Gunakan model yang sesuai, bukan Product

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Format data tidak valid",
		})
	}

	if err := h.service.CreateTransaction(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   req,
	})
}

// Update memperbarui data transaksi yang ada
func (h *TransactionHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	// 1. Parsing body langsung ke struct sementara
	var req models.Transaction
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Input tidak valid",
		})
	}

	// 2. Ambil data lama
	existing, err := h.service.GetTransactionByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Transaksi tidak ditemukan",
		})
	}

	existing.ProductID = req.ProductID
	existing.CustomerID = req.CustomerID

	if err := h.service.UpdateTransaction(existing); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal memperbarui transaksi",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   existing,
	})
}
