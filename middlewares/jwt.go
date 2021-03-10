package middlewares

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/vking34/fiber-messenger/utils"
)

// Protect protect routes
func Protect() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr, err := getJwtFromHeader(c, "authorization", "Bearer")
		if err != nil {
			return jwtError(c, err)
		}

		claims := jwt.MapClaims{}
		_, err = jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(utils.JwtSecret), nil
		})
		if err != nil {
			return jwtError(c, err)
		}

		c.Locals("userID", claims["userID"])
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

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
