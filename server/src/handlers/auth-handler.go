package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type SignInPayload struct {
	idToken string
}

func SignIn(c *fiber.Ctx) error {
	signInPayload := new(SignInPayload)

	if parsePayloadErr := c.BodyParser(signInPayload); parsePayloadErr != nil {
		log.Println(parsePayloadErr)
		c.Status(400).JSON(&fiber.Map{
			"message": parsePayloadErr,
		})

		return nil
	}

}
