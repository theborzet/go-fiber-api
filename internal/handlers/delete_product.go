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
	errchan := make(chan error)
	go func() {
		err := h.repo.DeleteProduct(productId)
		if err != nil {
			errchan <- fiber.NewError(fiber.StatusNotFound, err.Error())
		} else {
			errchan <- nil
		}
	}()

	err = <-errchan
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
