package middlewares

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/vking34/fiber-messenger/utils"
)

// Protect protect routes
func Protect() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, err := getJwtFromHeader(c, "authorization", "Bearer")
		if err != nil {
			return jwtError(c, err)
		}

		log.Println("token", token)
		return c.Next()
	}
}

func getJwtFromHeader(c *fiber.Ctx, header string, authScheme string) (string, error) {
	auth := c.Get(header)
	l := len(authScheme)
	if len(auth) > l+1 && auth[:l] == authScheme {
		return auth[l+1:], nil
	}
	return "", errors.New("Missing or malformed JWT")
}

func Protect1() fiber.Handler {
	return jwtware.New(jwtware.Config{
		ContextKey:   "token",
		SigningKey:   []byte(utils.JwtSecret),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
