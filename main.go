package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"log"
)

// quick start

func main() {
	// Initialize a new Fiber app
	app := fiber.New()
	app.Use(recover.New())

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello world, by fiber!")
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
