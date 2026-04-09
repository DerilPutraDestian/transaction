package routes

import (
	"transaction/handlers"
	"transaction/middleware" // Pastikan folder middleware sudah ada
	"transaction/repository"
	"transaction/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	userRepo := repository.NewAuthRepository(db)
	userSvc := service.NewAuthService(userRepo)
	userHandler := handlers.NewAuthHandler(userSvc)

	// Asset & Category
	productRepo := repository.NewProductRepository(db)
	productSvc := service.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productSvc)

	catRepo := repository.NewCategoryRepository(db)
	catSvc := service.NewCategoryService(catRepo)
	catHandler := handlers.NewCategoryHandler(catSvc)

	// Customer (DATA BARU)
	custRepo := repository.NewCustomerRepository(db)
	custSvc := service.NewCustomerService(custRepo)
	custHandler := handlers.NewCustomerHandler(custSvc)

	invoiceRepo := repository.NewInvoiceRepository(db)
	invoiceSvc := service.NewInvoiceService(invoiceRepo)
	invoiceHandler := handlers.NewInvoiceHandler(invoiceSvc)

	invoiceItemRepo := repository.NewInvoiceItemRepository(db)
	invoiceItemSvc := service.NewInvoiceItemService(invoiceItemRepo)
	invoiceItemHandler := handlers.NewInvoiceItemHandler(invoiceItemSvc)

	api := app.Group("/api")
	api.Post("/login", userHandler.Login)
	api.Use(middleware.JWTMiddleware())

	customers := api.Group("/customers")
	customers.Get("/", custHandler.Index)
	customers.Post("/", custHandler.Store)

	// --- PRODUCTS ---
	products := api.Group("/products")
	products.Get("/", productHandler.Index)
	products.Get("/:id", productHandler.Show)
	products.Post("/", productHandler.Store)    // Biasanya Admin
	products.Put("/:id", productHandler.Update) // Biasanya Admin
	products.Delete("/:id", productHandler.Delete)

	// --- CATEGORIES ---
	categories := api.Group("/categories")
	categories.Get("/", catHandler.Index)
	categories.Get("/:id", catHandler.Show)
	categories.Post("/", catHandler.Store)
	categories.Put("/:id", catHandler.Update)
	categories.Delete("/:id", catHandler.Delete)

	// --- INVOICES ---
	invoices := api.Group("/invoices")
	invoices.Get("/", invoiceHandler.Index)
	invoices.Get("/:id", invoiceHandler.Show)
	invoices.Post("/", invoiceHandler.Store)
	invoices.Put("/:id", invoiceHandler.Update)

	// --- INVOICE ITEMS ---
	invoiceItems := api.Group("/invoice-items")
	invoiceItems.Get("/", invoiceItemHandler.Index)
	invoiceItems.Get("/:id", invoiceItemHandler.Show)
	invoiceItems.Post("/", invoiceItemHandler.Store)
	invoiceItems.Put("/:id", invoiceItemHandler.Update)

}
