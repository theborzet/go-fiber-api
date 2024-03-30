package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	handler := &handler{
		DB: db,
	}

	routes := app.Group("/products")
	routes.Post("/", handler.AddProduct)
	routes.Delete("/:id", handler.DeleteProduct)
	routes.Get("/", handler.GetAllProduct)
	routes.Get("/:id", handler.GetProduct)
	routes.Put("/:id", handler.UpdateProduct)
}
