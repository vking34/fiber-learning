package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vking34/fiber-messenger/handlers"
)

// SetupRoutes setup route api
func SetupRoutes(app *fiber.App) {
	// Middlewares
	api := app.Group("/api/v1", logger.New())

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)
	auth.Post("/register", handlers.CreateUser)
}
