package main

import (
	Router "sudoku-api/src/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	Router.SetupRoutes(app)

	app.Listen(":3000")
}
