package general

import (
	"github.com/gofiber/fiber/v2"
)

func SetupKognitifRoutes(router fiber.Router) {
	kognitif := router.Group("/kognitif")

	kognitif.Get("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /kognitif")
	})

	kognitif.Get("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /kognitif/:id")
	})

	kognitif.Post("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("POST /kognitif")
	})

	kognitif.Put("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("PUT /kognitif/:id")
	})

	kognitif.Delete("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("DELETE /kognitif/:id")
	})
}
