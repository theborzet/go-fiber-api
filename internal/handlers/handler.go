package handlers

import (
	"go-fiber-api-docker/internal/db/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	repo *repository.SQLRepository
}

func NewHandler(repo *repository.SQLRepository) *Handler {
	return &Handler{repo: repo}
}

func RegisterRoutes(app *fiber.App, db *sqlx.DB) {
	handler := NewHandler(repository.NewSQLRepository(db))
	routes := app.Group("/products")
	routes.Post("/", handler.AddProduct)
	routes.Delete("/:id", handler.DeleteProduct)
	routes.Get("/", handler.GetAllProduct)
	routes.Get("/:id", handler.GetProduct)
	routes.Put("/:id", handler.UpdateProduct)
}
