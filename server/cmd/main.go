// main.go
package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/latoulicious/SIPP/internal/config"
	"github.com/latoulicious/SIPP/internal/cors"
	"github.com/latoulicious/SIPP/internal/database"
	"github.com/latoulicious/SIPP/internal/router"
	"github.com/latoulicious/SIPP/internal/util"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load configuration from .env file
	config.LoadConfig()

	// If JWT_SECRET is not set, generate a new JWT secret key
	if config.JwtSecret == "" {
		// Generate the JWT secret key
		secretKey, err := util.GenerateJWTSecretKey()
		if err != nil {
			logrus.Fatal("Error generating JWT secret key:", err)
		}

		// Print or log the generated JWT secret key
		logrus.Println("Generated JWT Secret Key:", secretKey)

		// Set the generated JWT secret key in the configuration
		config.JwtSecret = secretKey
	}

	// Configure Logrus logger to write logs to standard output
	logrus.SetOutput(os.Stdout)

	// Initialize Fiber app
	app := fiber.New()

	// Setup CORS
	cors.SetupCors(app)

	// Middleware to log HTTP method and path using Logrus
	app.Use(func(c *fiber.Ctx) error {
		logrus.Infof("%s %s", c.Method(), c.Path())
		return c.Next()
	})

	// Connect to the database
	err := database.ConnectDB()
	if err != nil {
		logrus.Fatal("Failed to connect to the database:", err)
	}

	// Setup routes
	router.SetupRoutes(app)

	// Start the server and handle errors
	err = app.Listen(":3000")
	if err != nil {
		logrus.Fatal("Failed to start the server:", err)
	}
}
