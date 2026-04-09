package handlers

import (
	"strconv"
	models "transaction/model"
	"transaction/service"

	"github.com/gofiber/fiber/v2"
)

type InvoiceHandler struct {
	service *service.InvoiceService
}

func NewInvoiceHandler(s *service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{service: s}
}

// Index mengambil daftar transaksi dengan paginasi
func (h *InvoiceHandler) Index(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	data, total, err := h.service.ListInvoices("", "", limit, offset)
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
func (h *InvoiceHandler) Store(c *fiber.Ctx) error {
	req := new(models.Invoice) // Gunakan model yang sesuai, bukan Product

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Format data tidak valid",
		})
	}

	if err := h.service.CreateInvoice(req); err != nil {
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

func (h *InvoiceHandler) Show(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := h.service.GetInvoiceByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Invoice tidak ditemukan",
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   data,
	})
}

// Update memperbarui data transaksi yang ada
func (h *InvoiceHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	// 1. Parsing body langsung ke struct sementara
	var req models.Invoice
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Input tidak valid",
		})
	}

	// 2. Ambil data lama
	existing, err := h.service.GetInvoiceByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Invoice tidak ditemukan",
		})
	}

	existing.CustomerID = req.CustomerID

	if err := h.service.UpdateInvoice(existing); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal memperbarui invoice",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   existing,
	})
}
