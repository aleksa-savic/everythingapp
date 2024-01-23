package main

import (
	database "everythingapp/db"
	"everythingapp/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// delete user in db by ID
func DeleteUserByID(c *fiber.Ctx) error {
	db := database.DB
	var user model.User
	// get id params
	id := c.Params("id")
	// find single user in the database by id
	db.Find(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	err := db.Delete(&user, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}
