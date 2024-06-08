package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		response := "HealthCheck OK!"

		return c.SendString(response)
	})

	app.Listen(":3000")
}
