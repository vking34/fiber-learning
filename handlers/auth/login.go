package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vking34/fiber-messenger/constants"
	"github.com/vking34/fiber-messenger/services/user_service"
	"github.com/vking34/fiber-messenger/utils"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login user login
func Login(c *fiber.Ctx) error {
	var req loginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Error on login request", "error": err})
	}

	username := req.Username
	pass := req.Password

	user, err := user_service.FindUser(username)
	if user == nil {
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": false, "message": "Can not login", "error": err})
		}

		return c.Status(400).JSON(&constants.InvalidUsernamePassword)
	}

	if !utils.CheckPasswordHash(pass, user.Password) {
		return c.Status(400).JSON(&constants.InvalidUsernamePassword)
	}

	token, err := utils.GenerateJWTToken(user.ID, username, 72)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": "Can not generate token",
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Logged in",
		"token":   token,
	})
}
