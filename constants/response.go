package constants

import (
	"github.com/gofiber/fiber/v2"
)

var UserNotFound fiber.Map
var InvalidUsernamePassword fiber.Map

func init() {
	UserNotFound = fiber.Map{
		"status":  false,
		"message": "User not found",
	}

	InvalidUsernamePassword = fiber.Map{
		"status":  false,
		"message": "Username/Password is invalid!",
	}
}
