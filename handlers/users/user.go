package users

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vking34/fiber-messenger/services/user_service"
)

// GetProfile return profile
func GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	log.Println("user id:", userID)

	user, err := user_service.FindUserByID(uint(userID.(float64)))
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "Can not get user",
			"error":   err,
		})
	}

	return c.JSON(user)
}
