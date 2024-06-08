package router

import (
	Handlers "sudoku-api/src/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// Middleware
	router := app.Group("/api", logger.New())

	Handlers.SetupUserHandlerRoutes(router)
	Handlers.SetupHealthcheckHandlerRoutes(router)
}
