package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// RegisterPayload
type RegisterPayload struct {
	Id            int    `json:"id"`
	Email         string `json:"email"`
	GoogleIdToken string `json:"googleIdToken"`
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

	log.Println(registerPayload)

	return c.JSON(registerPayload)
}

func SetupUserHandlerRoutes(router fiber.Router) {
	// Routes
	router.Post("/register", registerUser)
}
