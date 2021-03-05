package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/vking34/fiber-messenger/db"
	"github.com/vking34/fiber-messenger/routers"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	app := fiber.New()

	// middlewares
	app.Use(cors.New())
	app.Use(recover.New())

	// DB
	db.ConnectDB()

	// routers
	routers.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
