package general

import (
	"github.com/gofiber/fiber/v2"
)

func SetupSumatifRoutes(router fiber.Router) {
	sumatif := router.Group("/sumatif")

	sumatif.Get("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /sumatif")
	})

	sumatif.Get("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /sumatif/:id")
	})

	sumatif.Post("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("POST /sumatif")
	})

	sumatif.Put("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("PUT /sumatif/:id")
	})

	sumatif.Delete("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("DELETE /sumatif/:id")
	})
}
