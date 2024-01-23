package main

import (
	router "everythingapp/api/route"
	database "everythingapp/db"
	"fmt"
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
	router.SetupRoutes(app)

	app.Get("/", helloWorld)
	// Start the server
	log.Fatal(app.Listen(":3000"))

	fmt.Println("User created successfully!")

}
