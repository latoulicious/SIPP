// controller/user.go

package controller

import (
	"github.com/gofiber/fiber/v2"
	userHandler "github.com/latoulicious/SIPP/internal/handler/user"
)

func SetupUsersRoutes(router fiber.Router) {
	users := router.Group("/users")

	users.Get("/", userHandler.GetUsers)
	users.Get("/:usersId", userHandler.GetUsersById)
	users.Post("/", userHandler.CreateUsers)
	users.Put("/:usersId", userHandler.UpdateUsers)
	users.Delete("/:usersId", userHandler.DeleteUsers)
}
