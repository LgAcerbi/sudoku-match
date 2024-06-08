package router

import (
	Handlers "sudoku-api/src/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	router := app.Group("/api", logger.New())

	Handlers.SetupUserHandlerRoutes(router)
	Handlers.SetupHealthcheckHandlerRoutes(router)
}
