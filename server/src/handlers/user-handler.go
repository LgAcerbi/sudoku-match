package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	Services "sudoku-api/src/services"
)

type RegisterPayload struct {
	Email         string `json:"email"`
	GoogleIdToken string `json:"googleIdToken"`
	Nickname      string `json:"nickname"`
}

func registerUser(c *fiber.Ctx) error {
	registerPayload := new(RegisterPayload)

	if parsePayloadErr := c.BodyParser(registerPayload); parsePayloadErr != nil {
		log.Println(parsePayloadErr)
		c.Status(400).JSON(&fiber.Map{
			"message": parsePayloadErr,
		})
		return nil
	}

	registerData := Services.RegisterData{Email: registerPayload.Email, GoogleIdToken: registerPayload.GoogleIdToken, Nickname: registerPayload.Nickname}

	message, err := Services.RegisterUser(registerData)

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(&fiber.Map{
		"message": message,
	})
}

func SetupUserHandlerRoutes(router fiber.Router) {
	router.Post("/register", registerUser)
}
