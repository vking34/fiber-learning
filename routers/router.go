package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	authRouter "github.com/vking34/fiber-messenger/handlers/auth"
)

// SetupRoutes setup route api
func SetupRoutes(app *fiber.App) {
	// Middlewares
	api := app.Group("/api/v1", logger.New())

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", authRouter.Login)
	auth.Post("/register", authRouter.CreateUser)

	callbacks := auth.Group("/callbacks")
	callbacks.Post("/facebook", authRouter.FacebookCallback)
}
