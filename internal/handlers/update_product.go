package handlers

import (
	"go-fiber-api-docker/internal/db/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UpdateProductBody struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int64   `json:"stock"`
}

func (h Handler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	body := UpdateProductBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	product := models.Product{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Stock:       body.Stock,
	}

	errchan := make(chan error)

	go func() {
		err := h.repo.UpdateProduct(productId, &product)
		if err != nil {
			errchan <- fiber.NewError(fiber.StatusBadRequest, err.Error())
		} else {
			errchan <- nil
		}
	}()

	err = <-errchan
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(&product)
}
