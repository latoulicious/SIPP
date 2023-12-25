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
	capaianRepository := repository.NewCapaianRepository(database.DB) // Assuming you have capaianRepository
	capaianService := service.NewCapaianService(capaianRepository)
	capaianHandler := handler.NewCapaianHandler(capaianService)

	// Use UserService for AuthHandler
	authHandler := handler.NewAuthHandler(userService)

	// User routes with authentication and RBAC middleware
	userRoutes := api.Group("/user")
	userRoutes.Use(middleware.AuthMiddleware(e, userRepository))
	userRoutes.Use(middleware.RBACMiddleware(e, userRepository))
	userRoutes.Get("/", userHandler.GetUsers)
	userRoutes.Get("/:id", userHandler.GetUser)
	userRoutes.Post("/", userHandler.CreateUser)
	userRoutes.Put("/:id", userHandler.UpdateUser)
	userRoutes.Delete("/:id", userHandler.DeleteUser)

	// Add the new route for password change
	userRoutes.Post("/:id/change-password", userHandler.ChangePassword)

	// Authentication routes
	authRoutes := api.Group("/auth")
	authRoutes.Post("/login", authHandler.LoginHandler)

	// Capaian routes
	capaianRoutes := api.Group("/capaian")
	capaianRoutes.Get("/", capaianHandler.GetCapaian)
	capaianRoutes.Get("/:id", capaianHandler.GetCapaianByID)
	capaianRoutes.Post("/", capaianHandler.CreateCapaian)
	capaianRoutes.Put("/:id", capaianHandler.UpdateCapaian)
	capaianRoutes.Delete("/:id", capaianHandler.DeleteCapaian)

	// Add a new route in router/router.go
	app.Get("/api/get-jwt-secret", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"jwtSecret": config.JwtSecret})
	})

}
