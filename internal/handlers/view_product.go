package handlers

import (
	"encoding/json"
	"go-fiber-api-docker/internal/db/models"
	"go-fiber-api-docker/pkg/common/redis"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	cachedJSON, err := redis.RedisClient.Get("product:" + id).Result()
	if err == nil {
		var cachedProduct models.Product
		if result := json.Unmarshal([]byte(cachedJSON), &cachedProduct); result != nil {
			return fiber.NewError(fiber.StatusNoContent, result.Error())
		}

		return c.Status(fiber.StatusOK).JSON(&cachedProduct)
	}

	var product *models.Product

	productId, _ := strconv.Atoi(id)

	errchan := make(chan error)

	go func() {
		product, err = h.repo.GetProductID(productId)
		if err != nil {
			errchan <- fiber.NewError(fiber.StatusNotFound, err.Error())
		} else {
			errchan <- nil
		}
	}()

	//Преобразуем продукт в JSON
	productJSON, err := json.Marshal(product)
	if err != nil {
		return fiber.NewError(fiber.StatusNotImplemented, err.Error())
	}

	if err := redis.RedisClient.Set("product:"+id, productJSON, 0).Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Ошибка при сохранении в Redis",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&product)
}
