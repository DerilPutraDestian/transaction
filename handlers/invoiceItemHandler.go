package handlers

import (
	"strconv"
	models "transaction/model"
	"transaction/service"

	"github.com/gofiber/fiber/v2"
)

type InvoiceItemHandler struct {
	service *service.InvoiceItemService
}

func NewInvoiceItemHandler(s *service.InvoiceItemService) *InvoiceItemHandler {
	return &InvoiceItemHandler{service: s}
}

// Index mengambil daftar transaksi dengan paginasi
func (h *InvoiceItemHandler) Index(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	data, total, err := h.service.ListInvoiceItems("", "", limit, offset)
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
func (h *InvoiceItemHandler) Show(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := h.service.GetInvoiceItemByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Data tidak ditemukan",
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   data,
	})
}

// Store membuat transaksi baru
func (h *InvoiceItemHandler) Store(c *fiber.Ctx) error {
	req := new(models.InvoiceItem) // Gunakan model yang sesuai, bukan Product

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Format data tidak valid",
		})
	}

	if err := h.service.CreateInvoiceItem(req); err != nil {
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
func (h *InvoiceItemHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var req models.InvoiceItem
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Input tidak valid",
		})
	}
	existing, err := h.service.GetInvoiceItemByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Invoice item tidak ditemukan",
		})
	}

	existing.InvoiceID = req.InvoiceID
	existing.ProductID = req.ProductID
	existing.Qty = req.Qty
	existing.UnitPrice = req.UnitPrice

	if err := h.service.UpdateInvoiceItem(existing); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal memperbarui data",
		})
	}
	updatedData, _ := h.service.GetInvoiceItemByID(id)

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   updatedData,
	})
}
