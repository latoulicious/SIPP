package general

import (
	"github.com/gofiber/fiber/v2"
)

func SetupKesukaranRoutes(router fiber.Router) {
	kesukaran := router.Group("/kesukaran")

	kesukaran.Get("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /kesukaran")
	})

	kesukaran.Get("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /kesukaran/:id")
	})

	kesukaran.Post("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("POST /kesukaran")
	})

	kesukaran.Put("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("PUT /kesukaran/:id")
	})

	kesukaran.Delete("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("DELETE /kesukaran/:id")
	})
}
