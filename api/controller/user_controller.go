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

func UpdateUserById(c *fiber.Ctx) error {

	u := new(model.User)
	db := database.DB

	id := c.Params("id")
	//trazenjen po id
	db.Find(&u, "id = ?", id)
	if u.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	err := c.BodyParser(u)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	db.Save(&u)
	// Return the updated user
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "users Found", "data": u})
}

func UpdateUserByEmail(c *fiber.Ctx) error {

	u := new(model.User)
	db := database.DB

	email := c.Params("email")

	result := db.Where("email = ?", email).First(&u)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	// Parsiranje za ažuriranje podataka korisnika
	err := c.BodyParser(u)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	result = db.Save(&u)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to update user", "data": result.Error})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User updated", "data": u})
}
