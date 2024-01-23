package main

import (
	database "everythingapp/db"
	"log"

	"everythingapp/api/route"

	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
func initDatabase() {

}
func main() {
	app := fiber.New()

	route.SetupRoutes(app)

	database.ConnectDB()
	app.Get("/", helloWorld)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	// Start the server
	log.Fatal(app.Listen(":3000"))
}
