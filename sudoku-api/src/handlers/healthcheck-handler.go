package handlers

import "github.com/gofiber/fiber/v2"

func healthcheck(c *fiber.Ctx) error {
	response := "HealthCheck OK!"

	return c.SendString(response)
}

func SetupHealthcheckHandlerRoutes(router fiber.Router) {
	router.Get("/healthcheck", healthcheck)
}
