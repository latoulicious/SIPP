package general

import (
	"github.com/gofiber/fiber/v2"
)

func SetupPengayaanRoutes(router fiber.Router) {
	pengayaan := router.Group("/pengayaan")

	pengayaan.Get("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /pengayaan")
	})

	pengayaan.Get("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /pengayaan/:id")
	})

	pengayaan.Post("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("POST /pengayaan")
	})

	pengayaan.Put("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("PUT /pengayaan/:id")
	})

	pengayaan.Delete("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("DELETE /pengayaan/:id")
	})
}
