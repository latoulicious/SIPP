package general

import (
	"github.com/gofiber/fiber/v2"
)

func SetupFormatifRoutes(router fiber.Router) {
	formatif := router.Group("/formatif")

	formatif.Get("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /formatif")
	})

	formatif.Get("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /formatif/:id")
	})

	formatif.Post("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("POST /formatif")
	})

	formatif.Put("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("PUT /formatif/:id")
	})

	formatif.Delete("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("DELETE /formatif/:id")
	})
}
