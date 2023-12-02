package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/latoulicious/SIPP/internal/database"
	"github.com/latoulicious/SIPP/internal/router"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Connect to the database
	database.ConnectDB()

	// Setup routes
	router.SetupRoutes(app)

	// Start the server and handle errors
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
