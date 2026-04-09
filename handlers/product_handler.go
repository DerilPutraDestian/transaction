package handlers

import (
	"strconv"
	models "transaction/model"
	"transaction/service"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) Index(c *fiber.Ctx) error {
	categoryID := c.Query("category_id")
	search := c.Query("search")
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	products, count, err := h.service.ListProducts(categoryID, search, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"count":  count,
		"data":   products,
	})
}

// GET /api/products/:id
func (h *ProductHandler) Show(c *fiber.Ctx) error {
	id := c.Params("id")

	product, err := h.service.GetProduct(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Product not found"})
	}

	return c.JSON(fiber.Map{"status": "success", "data": product})
}

// POST /api/products
func (h *ProductHandler) Store(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid request body"})
	}
	if err := h.service.CreateProduct(&product); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "data": product})
}

// PUT /api/products/:id
func (h *ProductHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	// 1. Cari data lama
	product, err := h.service.GetProduct(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Asset not found"})
	}

	// 2. Timpa dengan data baru dari Body
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	// 3. Simpan perubahan
	if err := h.service.UpdateProduct(product); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": product})
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	product, err := h.service.GetProduct(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Asset not found"})
	}

	if err := h.service.DeleteProduct(product); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Product deleted"})
}
