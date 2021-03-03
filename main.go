package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/vking34/fiber-messenger/routers"
)

func main() {
	app := fiber.New()

	// middlewares
	app.Use(cors.New())
	app.Use(recover.New())

	// routers
	routers.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
