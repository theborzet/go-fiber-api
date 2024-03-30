package handlers

import (
	"go-fiber-api-docker/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

type UpdateProductBody struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int64   `json:"stock"`
}

func (h *handler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateProductBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}

	product.Name = body.Name
	product.Description = body.Description
	product.Prise = body.Price
	product.Stock = body.Stock

	h.DB.Save(&product)

	return c.Status(fiber.StatusCreated).JSON(&product)
}
