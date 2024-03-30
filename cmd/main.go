package main

import (
	"go-fiber-api-docker/pkg/common/config"
	"go-fiber-api-docker/pkg/common/db"
	"go-fiber-api-docker/pkg/common/redis"
	"go-fiber-api-docker/pkg/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	app := fiber.New()

	h := db.Init(c)

	handlers.RegisterRoutes(app, h)
	redis.SetupClient()

	app.Listen(c.Port)
}
