package main

import "github.com/gofiber/fiber/v2"

func registerRoutes(app *fiber.App) {
	userGroup := app.Group("/users")
	{
		userGroup.Post("/", create)
		userGroup.Get("/", read)
		userGroup.Get("/:userId", getUser)
		userGroup.Put("/:userId", update)
		userGroup.Patch("/:userId", partialupdate)
		userGroup.Delete("/:userId", delete)
	}
}
