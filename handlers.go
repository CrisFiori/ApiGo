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
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	users = append(users, userFromBody)
	return c.JSON(userFromBody)

}

func read(c *fiber.Ctx) error {
	return c.JSON(users)
}

func getUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	userFound, _, err := findUserById(userId)
	if err != nil {
		return c.JSON(err.Error())
	}
	return c.JSON(userFound)
}

func update(c *fiber.Ctx) error {
	userId := c.Params("userID")
	var userFromBody userType

	_, index, err := findUserById(userId)
	if err != nil {
		return c.JSON(err.Error())
	}
	c.BodyParser(&userFromBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	//replace value from users
	//	first: remove from users, second add to users

	//removing users
	users = append(users[:index], users[index+1:]...)
	//adding users
	users = append(users, userFromBody)
	return c.SendString("User updated")
}

func partialupdate(c *fiber.Ctx) error {
	userId := c.Params("userID")
	userFound, _, err := findUserById(userId)
	if err != nil {
		return c.JSON(err.Error())
	}
	//Se valido todo correctamente y continuamos con el codigo
	userFound.IsAdmin = !userFound.IsAdmin
	return c.SendString("Is Admin updated")
}

func delete(c *fiber.Ctx) error {
	userId := c.Params("userID")
	_, index, err := findUserById(userId)
	if err != nil {
		return c.JSON(err.Error())
	}
	//remove with slice operator
	users = append(users[:index], users[index+1:]...)
	return c.SendString("User Removed")
}

func findUserById(userId string) (*userType, int, error) {
	//loop over users
	for index, user := range users {
		if user.Id == userId {
			return &users[index], index, nil
		}
	}
	return nil, -1, errors.New("user not found")
}
