package route

import (
	"everythingapp/api/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// grouping

	api := app.Group("/api")
	v1 := api.Group("/user") // routes
	v1.Get("/", controller.GetAllUsers)

}
