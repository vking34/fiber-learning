package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	authHandlers "github.com/vking34/fiber-messenger/handlers/auth"
	userHandlers "github.com/vking34/fiber-messenger/handlers/users"
	"github.com/vking34/fiber-messenger/middlewares"
)

// SetupRoutes setup route api
func SetupRoutes(app *fiber.App) {
	// Middlewares
	api := app.Group("/api/v1", logger.New())

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", authHandlers.Login)
	auth.Post("/register", authHandlers.CreateUser)

	callbacks := auth.Group("/callbacks")
	callbacks.Post("/facebook", authHandlers.FacebookCallback)

	// User
	user := api.Group("/users")
	user.Get("/me", middlewares.Protect(), userHandlers.GetProfile)
}
