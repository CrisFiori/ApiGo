package main

import (
	"github.com/gofiber/fiber/v2"
)

type userType struct {
	Id      string
	Name    string
	IsAdmin bool
}

var users = []userType{
	{Id: "1", Name: "Iron Man", IsAdmin: true},
	{Id: "2", Name: "Thor", IsAdmin: true},
	{Id: "3", Name: "Spiderman", IsAdmin: true},
}

func main() {
	app := fiber.New()

	//routes GET
	//CRUD: create -read -update -delete
	userGroup := app.Group("/users")
	{
		userGroup.Post("/", create)
		userGroup.Get("/", read)
		userGroup.Get("/:userId", getUser)
		userGroup.Put("/:userId", update)
		userGroup.Patch("/:userId", partialupdate)
		userGroup.Delete("/:userId", delete)
	}

	//Start Server
	app.Listen(":8080")
}
