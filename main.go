package main

import (
	route "everythingapp/api/route"
	database "everythingapp/db"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func main() {
	app := fiber.New()

	database.ConnectDB()
	route.SetupRoutes(app)
	app.Get("/", helloWorld)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	// Start the server
	log.Fatal(app.Listen(":3000"))

	fmt.Println("User created successfully!")
}
