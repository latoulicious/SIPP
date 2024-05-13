package general

import (
	"github.com/gofiber/fiber/v2"
)

func SetupTingkatRoutes(router fiber.Router) {
	tingkat := router.Group("/tingkat")

	tingkat.Get("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /tingkat")
	})

	tingkat.Get("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /tingkat/:id")
	})

	tingkat.Post("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("POST /tingkat")
	})

}
