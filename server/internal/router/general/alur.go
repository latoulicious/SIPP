package general

import (
	"github.com/gofiber/fiber/v2"
)

func SetupAlurRoutes(router fiber.Router) {
	alur := router.Group("/alur")

	alur.Get("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /alur")
	})

	alur.Get("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /alur/:id")
	})

	alur.Post("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("POST /alur")
	})

	alur.Put("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("PUT /alur/:id")
	})

	alur.Delete("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("DELETE /alur/:id")
	})
}
