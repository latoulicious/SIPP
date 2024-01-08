package general

import (
	"github.com/gofiber/fiber/v2"
)

func SetupModulRoutes(router fiber.Router) {
	modul := router.Group("/modul")

	modul.Get("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /modul")
	})

	modul.Get("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /modul/:id")
	})

	modul.Post("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("POST /modul")
	})

	modul.Put("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("PUT /modul/:id")
	})

	modul.Delete("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("DELETE /modul/:id")
	})
}
