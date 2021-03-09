package users

import (
	"github.com/gofiber/fiber/v2"
)

// GetProfile return profile
func GetProfile(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"user_id": "1",
	})
}
