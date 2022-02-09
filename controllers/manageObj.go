package controllers

import (
	"TaskManager/database"
	"TaskManager/repo"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllObject(c *fiber.Ctx) error {
	return c.JSON(database.DB.GetAllObj())
}

func GetObj(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	return c.JSON(database.DB.GetObj(id))
}

/*
func GetObjWithCom(c *fiber.Ctx) error {
	idO, _ := strconv.Atoi(c.Params("id"))
	idC, _ := strconv.Atoi(c.Params("num"))
	return c.JSON(database.DB.GetObjWithCom(idO, idC))
}
*/
func CreateObj(c *fiber.Ctx) error {

	var object = repo.Object{}

	err := c.BodyParser(&object)
	if err != nil {
		return err
	}

	return c.JSON(database.DB.CreateObj(&object))
}
func UpdateObj(c *fiber.Ctx) error {

	var object = repo.Object{}
	id, _ := strconv.Atoi(c.Params("id"))
	err := c.BodyParser(&object)
	if err != nil {
		return err
	}

	return c.JSON(database.DB.UpdateObj(&object, id))
}

func DeleteObj(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return err
	}

	return c.JSON(database.DB.DeleteObj(id))
}
