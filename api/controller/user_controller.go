package controller

import (
	database "everythingapp/db"
	"everythingapp/model"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has created", "data": user})
}
