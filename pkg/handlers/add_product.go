package handlers

import (
	"go-fiber-api-docker/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

type AddProductBody struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int64   `json:"stock"`
}

func (h *handler) AddProduct(c *fiber.Ctx) error {
	body := AddProductBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var product models.Product

	product.Name = body.Name
	product.Description = body.Description
	product.Prise = body.Price
	product.Stock = body.Stock

	if result := h.DB.Create(&product); result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&product)
}
