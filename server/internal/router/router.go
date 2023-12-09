// router/router.go

package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/latoulicious/SIPP/internal/config"
	"github.com/latoulicious/SIPP/internal/database"
	"github.com/latoulicious/SIPP/internal/handler"
	"github.com/latoulicious/SIPP/internal/middleware"
	"github.com/latoulicious/SIPP/internal/repository"
	"github.com/latoulicious/SIPP/internal/service"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New())

	api := app.Group("/api", logger.New())

	// Initialize repositories and services
	userRepository := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// Use UserService for AuthHandler
	authHandler := handler.NewAuthHandler(userService) // Change this line

	// User routes with authentication middleware
	userRoutes := api.Group("/users")
	userRoutes.Use(middleware.AuthMiddleware)
	userRoutes.Get("/", userHandler.GetUsers)
	userRoutes.Get("/:id", userHandler.GetUser)
	userRoutes.Post("/", userHandler.CreateUser)
	userRoutes.Put("/:id", userHandler.UpdateUser)
	userRoutes.Delete("/:id", userHandler.DeleteUser)

	// Authentication routes
	authRoutes := api.Group("/auth")
	authRoutes.Post("/login", authHandler.LoginHandler) // Assuming you have a login handler

	// More authentication routes can be added here if needed

	// Add a new route in router/router.go
	app.Get("/api/get-jwt-secret", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"jwtSecret": config.JwtSecret})
	})
}
