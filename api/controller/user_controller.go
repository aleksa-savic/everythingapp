package controller

import (
	database "everythingapp/db"
	"everythingapp/model"
	"everythingapp/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func DeleteUserByID(c *fiber.Ctx) error {
	db := database.DB
	var user model.User
	id := c.Params("id")
	db.Find(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	err := db.Unscoped().Delete(&user, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	hashedPassword, hash_err := utils.HashPassword(user.Password)
	if hash_err != nil {
		return hash_err
	}

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	user.Password = hashedPassword

	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has created", "data": user})

}

func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []model.User
	db.Find(&users)
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Users Found", "data": users})
}
