package handlers

import "github.com/gofiber/fiber/v2"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Error on login request", "error": err})
	}

	username := req.Username
	pass := req.Password

	return c.JSON(fiber.Map{
		"status":   true,
		"username": username,
		"pass":     pass,
	})
}
