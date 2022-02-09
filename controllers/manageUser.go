package controllers

import (
	"TaskManager/database"
	"fmt"

	"TaskManager/repo"

	"github.com/gofiber/fiber/v2"
)

func AllUsers(c *fiber.Ctx) error {
	return c.JSON(database.DB.GetAllUsers())
}

func CreaterUser(c *fiber.Ctx) error {
	var user = repo.User{}

	err := c.BodyParser(&user)
	if err != nil {
		return err
	}

	fmt.Println(user)
	user.CreatePassword("Love")
	database.DB.CreateUser(&user)

	return c.JSON(user)
}
