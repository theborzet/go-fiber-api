package handlers

import (
	"encoding/json"
	"go-fiber-api-docker/pkg/common/models"
	"go-fiber-api-docker/pkg/common/redis"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetAllProduct(c *fiber.Ctx) error {
	keys, err := redis.RedisClient.Keys("product:*").Result()
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}

	var products []models.Product

	for _, key := range keys {
		cachedJSON, err := redis.RedisClient.Get("product:" + key).Result()
		if err != nil {
			continue
		}

		var cachedProduct models.Product
		if result := json.Unmarshal([]byte(cachedJSON), &cachedProduct); result != nil {
			continue
		}
		products = append(products, cachedProduct)

	}
	if len(products) == 0 {
		if result := h.DB.Find(&products); result.Error != nil {
			return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
		}
		for _, product := range products {
			productJSON, err := json.Marshal(product)
			if err != nil {
				continue
			}
			if err := redis.RedisClient.Set("product:"+strconv.Itoa(product.ID), productJSON, 0).Err(); err != nil {
				continue
			}
		}
	}

	return c.Status(fiber.StatusOK).JSON(&products)
}
