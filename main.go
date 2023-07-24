package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	//middlewares
	app.Use(logger.New())
	//routes
	registerRoutes(app)

	//Start Server
	app.Listen(":8080")
}
