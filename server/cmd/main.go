package main

import (
	"fmt"
	"log"
	"os"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/latoulicious/SIPP/internal/config"
	"github.com/latoulicious/SIPP/internal/cors"
	"github.com/latoulicious/SIPP/internal/database"
	"github.com/latoulicious/SIPP/internal/router"
	"github.com/latoulicious/SIPP/internal/util"
)

// Declare dsn at the package level
var DSN string

// main is the entry point for the application. It loads configuration, initializes dependencies like the database, sets up routes and middleware, and starts the HTTP server.
func main() {
	// Load configuration from .env file
	config.LoadConfig()

	if os.Getenv("ENVIRONMENT") == "development" {
		// If JWT_SECRET is not set, generate a new JWT secret key
		if config.JwtSecret == "" {
			// Generate the JWT secret key
			secretKey, err := util.GenerateJWTSecretKey()
			if err != nil {
				log.Fatal("Error generating JWT secret key:", err)
			}

			// Print or log the generated JWT secret key
			log.Println("Generated JWT Secret Key:", secretKey)

			// Set the generated JWT secret key in the configuration
			config.JwtSecret = secretKey
		}
	}

	// Initialize Fiber app
	app := fiber.New()

	// Setup CORS
	cors.SetupCors(app)

	// Logging Request ID
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	// Connect to the database
	err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Somewhere in your application startup code
	enforcer, err := initCasbin()
	if err != nil {
		// Handle initialization error
		log.Fatal(err)
	}

	// Setup routes
	router.SetupRoutes(app, enforcer)

	// Start the server and handle errors
	err = app.Listen(":3000")
	if err != nil {
		log.Fatal("Failed to start the server:", err)
	}

	// Graceful shutdown
	err = app.Shutdown()
	if err != nil {
		log.Println("Graceful shutdown failed:", err)
	}
}

// initCasbin initializes a Casbin enforcer using the provided model file and policy file.
// It returns a Casbin enforcer instance to be used for access control.

func initCasbin() (*casbin.Enforcer, error) {
	// Load the model configuration from the file
	m, err := model.NewModelFromFile("./internal/config/casbin_model.conf")
	if err != nil {
		return nil, fmt.Errorf("error loading Casbin model: %v", err)
	}

	// Create a file adapter for policy storage
	a := fileadapter.NewAdapter("./internal/config/casbin_policy.csv")

	// Log information about model and adapter
	log.Println("Casbin model and policy adapter initialized successfully")

	// Create a Casbin enforcer with the model and adapter
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, fmt.Errorf("error initializing Casbin enforcer: %v", err)
	}

	// Enable logging for debugging purposes
	e.EnableLog(false)

	log.Println("Casbin enforcer initialized successfully")

	// After initializing Casbin enforcer
	err = e.LoadPolicy()
	if err != nil {
		return nil, fmt.Errorf("error loading Casbin policies: %v", err)
	}

	log.Println("Casbin policies loaded successfully")

	return e, nil
}
