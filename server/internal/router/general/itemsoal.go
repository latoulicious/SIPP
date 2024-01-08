package general

import (
	"github.com/gofiber/fiber/v2"
)

func SetupItemSoalRoutes(router fiber.Router) {
	itemsoal := router.Group("/itemsoal")

	itemsoal.Get("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /itemsoal")
	})

	itemsoal.Get("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /itemsoal/:id")
	})

	itemsoal.Post("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("POST /itemsoal")
	})

	itemsoal.Put("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("PUT /itemsoal/:id")
	})

	itemsoal.Delete("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("DELETE /itemsoal/:id")
	})
}
