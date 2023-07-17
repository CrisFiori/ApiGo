package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func create(c *fiber.Ctx) error {
	//add new user to users
	var userFromBody userType
	err := c.BodyParser(&userFromBody)

	if err != nil {
		return c.SendString("Bad request")
	}
	users = append(users, userFromBody)
	return c.JSON(userFromBody)

}

func read(c *fiber.Ctx) error {
	return c.JSON(users)
}

func getUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	userFound, err := findUserById(userId)
	if err != nil {
		return c.JSON(err.Error())
	}
	return c.JSON(userFound)
}

func update(c *fiber.Ctx) error {
	return c.SendString("Hello Fiber")
}

func delete(c *fiber.Ctx) error {
	return c.SendString("Hello Fiber")
}

func findUserById(userId string) (*userType, error) {
	//loop over users
	for index, user := range users {
		if user.Id == userId {
			return &users[index], nil
		}
	}
	return nil, errors.New("user not found")
}
