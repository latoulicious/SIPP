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
	mapelRepository := repository.NewMapelRepository(database.DB)
	mapelService := service.NewMapelService(mapelRepository)
	mapelHandler := handler.NewMapelHandler(mapelService)
	kelasRepository := repository.NewKelasRepository(database.DB)
	kelasService := service.NewKelasService(kelasRepository)
	kelasHandler := handler.NewKelasHandler(kelasService)
	tahunRepository := repository.NewTahunRepository(database.DB)
	tahunService := service.NewTahunService(tahunRepository)
	tahunHandler := handler.NewTahunHandler(tahunService)
	capaianRepository := repository.NewCapaianRepository(database.DB)
	capaianService := service.NewCapaianService(capaianRepository)
	capaianHandler := handler.NewCapaianHandler(capaianService)

	// Use UserService for AuthHandler
	authHandler := handler.NewAuthHandler(userService)

	// Core routess

	// User routes with authentication and RBAC middleware
	userRoutes := api.Group("/user")
	userRoutes.Use(middleware.AuthMiddleware(e, userRepository))
	userRoutes.Use(middleware.RBACMiddleware(e, userRepository))
	userRoutes.Get("/", userHandler.GetUsers)
	userRoutes.Get("/:id", userHandler.GetUserByID)
	userRoutes.Post("/", userHandler.CreateUser)
	userRoutes.Put("/:id", userHandler.UpdateUser)
	userRoutes.Delete("/:id", userHandler.DeleteUser)

	// Add the new route for password change
	userRoutes.Post("/:id/change-password", userHandler.ChangePassword)

	// Authentication routes
	authRoutes := api.Group("/auth")
	authRoutes.Post("/login", authHandler.LoginHandler)

	// Mapel routes
	mapelRoutes := api.Group("/mapel")
	mapelRoutes.Get("/", mapelHandler.GetMapel)
	mapelRoutes.Get("/:id", mapelHandler.GetMapelByID)
	mapelRoutes.Post("/", mapelHandler.CreateMapel)
	mapelRoutes.Put("/:id", mapelHandler.UpdateMapel)
	mapelRoutes.Delete("/:id", mapelHandler.DeleteMapel)

	// Kelas routes
	kelasRoutes := api.Group("/kelas")
	kelasRoutes.Get("/", kelasHandler.GetKelas)
	kelasRoutes.Get("/:id", kelasHandler.GetKelasByID)
	kelasRoutes.Post("/", kelasHandler.CreateKelas)
	kelasRoutes.Put("/:id", kelasHandler.UpdateKelas)
	kelasRoutes.Delete("/:id", kelasHandler.DeleteKelas)

	// Tahun routes
	tahunRoutes := api.Group("/tahun")
	tahunRoutes.Get("/", tahunHandler.GetTahun)
	tahunRoutes.Get("/:id", tahunHandler.GetTahunByID)
	tahunRoutes.Post("/", tahunHandler.CreateTahun)
	tahunRoutes.Put("/:id", tahunHandler.UpdateTahun)
	tahunRoutes.Delete("/:id", tahunHandler.DeleteTahun)

	// General usage routes

	// Capaian routes
	capaianRoutes := api.Group("/capaian")
	capaianRoutes.Get("/", capaianHandler.GetCapaian)
	capaianRoutes.Get("/:id", capaianHandler.GetCapaianByID)
	capaianRoutes.Post("/", capaianHandler.CreateCapaian)
	capaianRoutes.Put("/:id", capaianHandler.UpdateCapaian)
	capaianRoutes.Delete("/:id", capaianHandler.DeleteCapaian)

	// misc routes
	app.Get("/api/get-jwt-secret", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"jwtSecret": config.JwtSecret})
	})

}
