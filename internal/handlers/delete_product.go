package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if err := h.repo.DeleteProduct(productId); err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
