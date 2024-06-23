package handlers

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	Services "sudoku-api/src/services"
)

type RegisterPayload struct {
	Email         string `json:"email"`
	GoogleIdToken string `json:"googleIdToken"`
	Nickname      string `json:"nickname"`
}

type UserOutput struct {
	Email     string     `json:"email"`
	Nickname  string     `json:"nickname"`
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func RegisterUser(c *fiber.Ctx) error {
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

func FindUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")

	foundUser, err := Services.FindUserByEmail(email)

	if err != nil {
		return c.Status(404).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	response := UserOutput{
		Nickname:  foundUser.Nickname,
		Email:     foundUser.Email,
		ID:        foundUser.ID,
		CreatedAt: foundUser.CreatedAt,
		UpdatedAt: foundUser.UpdatedAt,
		DeletedAt: foundUser.DeletedAt,
	}

	return c.Status(200).JSON(response)
}

func SetupUserHandlerRoutes(router fiber.Router) {
	router.Post("/register", RegisterUser)
	router.Get("/user/:email", FindUserByEmail)
}
