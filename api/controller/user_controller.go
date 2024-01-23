package controller

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
	err := db.Unscoped().Delete(&user, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}
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

func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []model.User // find all users in the database
	db.Find(&users)        // If no user found, return an error
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	} // return users
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Users Found", "data": users})
}

func GetSingleUserByID(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var user model.User
	db.Find(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found by id", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
}

func GetSingleUserByUsername(c *fiber.Ctx) error {
	db := database.DB
	username := c.Params("username")
	var user model.User
	db.Find(&user, "username = ?", username)
	if user.Username == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found by username", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
}
