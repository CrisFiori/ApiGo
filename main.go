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
	{Id: "1", Name: "Georgi", IsAdmin: true},
	{Id: "2", Name: "Cristian", IsAdmin: true},
	{Id: "3", Name: "Delfi", IsAdmin: true},
}

func main() {
	app := fiber.New()

	//routes GET
	//CRUD: create -read -update -delete

	app.Post("/", create)
	app.Get("/users", read)
	app.Get("/users/:userId", getUser)
	app.Put("/", update)
	app.Patch("/", update)
	app.Delete("/", delete)
	//Start Server
	app.Listen(":8080")
}
