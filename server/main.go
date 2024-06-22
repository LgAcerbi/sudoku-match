package main

import (
	"os"
	Router "sudoku-api/src/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := fiber.New()

	Router.SetupRoutes(app)

	app.Listen("127.0.0.1:" + os.Getenv("PORT"))
}
