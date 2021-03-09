package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vking34/fiber-messenger/models"
	"github.com/vking34/fiber-messenger/utils"
)

// GenerateToken gen token
func GenerateToken(c *fiber.Ctx, user *models.User) error {
	token, err := utils.GenerateJWTToken(user.ID, user.Username, 72)
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
