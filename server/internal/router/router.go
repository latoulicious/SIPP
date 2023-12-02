package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	userController "github.com/latoulicious/SIPP/internal/controller/user"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	userController.SetupUsersRoutes(api)
}
