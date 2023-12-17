// router/router.go

package router

import (
	"github.com/casbin/casbin/v2"
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

func SetupRoutes(app *fiber.App, e *casbin.Enforcer) {
	app.Use(cors.New())

	api := app.Group("/api", logger.New())

	// Initialize repositories and services
	userRepository := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// Use UserService for AuthHandler
	authHandler := handler.NewAuthHandler(userService) // Change this line

	// User routes with authentication and RBAC middleware
	userRoutes := api.Group("/users")
	userRoutes.Use(middleware.AuthMiddleware(e, userRepository))
	userRoutes.Use(middleware.RBACMiddleware(e, userRepository)) // Keep the second argument
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

	// Add a new route to get the user's role based on the JWT token
	app.Get("/api/get-user-role", middleware.AuthMiddleware(e, userRepository), func(c *fiber.Ctx) error {
		// Retrieve the user's role from the context (set by AuthMiddleware)
		userRole, ok := c.Locals("user_role").(string)
		if !ok {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve user role"})
		}

		return c.JSON(fiber.Map{"role": userRole})
	})

}
