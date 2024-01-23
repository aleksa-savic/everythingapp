package route

import (
	"everythingapp/api/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/user")

	v1.Post("/", controller.CreateUser)
	v1.Get("/", controller.GetAllUsers)
	v1.Delete("/:id", controller.DeleteUserByID)
	v1.Put(":id", controller.UpdateUserById)
	v1.Put("/email/:email", controller.UpdateUserByEmail)
}
