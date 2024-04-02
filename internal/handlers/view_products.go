package handlers

import (
	"encoding/json"
	"go-fiber-api-docker/internal/db/models"
	"go-fiber-api-docker/pkg/common/redis"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetAllProduct(c *fiber.Ctx) error {
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
		results, err := h.repo.GetAllProduct()
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		for _, product := range results {
			productJSON, err := json.Marshal(product)
			if err != nil {
				continue
			}
			if err := redis.RedisClient.Set("product:"+strconv.Itoa(product.ID), productJSON, 0).Err(); err != nil {
				continue
			}
		}
		return c.Status(fiber.StatusOK).JSON(&results)
	}

	return c.Status(fiber.StatusOK).JSON(&products)
}
