package general

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRemedialRoutes(router fiber.Router) {
	remedial := router.Group("/remedial")

	remedial.Get("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /remedial")
	})

	remedial.Get("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /remedial/:id")
	})

	remedial.Post("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("POST /remedial")
	})

	remedial.Put("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("PUT /remedial/:id")
	})

	remedial.Delete("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("DELETE /remedial/:id")
	})
}
