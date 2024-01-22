package main

import (
	database "everythingapp/db"
	"log"

	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
func initDatabase() {

}
func main() {
	app := fiber.New()

	database.ConnectDB()
	app.Get("/", helloWorld)
	// Start the server
	log.Fatal(app.Listen(":3000"))
}
