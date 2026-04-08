package handlers

import (
	"strconv"
	models "transaction/model"
	"transaction/service"

	"github.com/gofiber/fiber/v2"
)

type customerHandler struct {
	service *service.CustomerService
}

func NewCustomerHandler(s *service.CustomerService) *customerHandler {
	return &customerHandler{service: s}
}

// GET /api/customers
func (h *customerHandler) Index(c *fiber.Ctx) error {
	search := c.Query("search")
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	data, total, err := h.service.ListCustomers(search, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"count":  total,
		"data":   data,
	})
}

// POST /api/customers
func (h *customerHandler) Store(c *fiber.Ctx) error {
	var emp models.Customer

	if err := c.BodyParser(&emp); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid request body"})
	}

	if err := h.service.CreateCustomer(&emp); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "data": emp})
}
