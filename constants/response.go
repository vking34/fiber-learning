package constants

import (
	"github.com/gofiber/fiber/v2"
)

var UserNotFound fiber.Map
var InvalidUsernamePassword fiber.Map
var ExistingUsername fiber.Map
var ExistingEmail fiber.Map

func init() {
	UserNotFound = fiber.Map{
		"status":  false,
		"message": "User not found",
	}

	InvalidUsernamePassword = fiber.Map{
		"status":  false,
		"message": "Username/Password is invalid!",
	}

	ExistingUsername = fiber.Map{
		"status":  false,
		"message": "Existing Username",
	}

	ExistingEmail = fiber.Map{
		"status":  false,
		"message": "Existing Email",
	}
}
