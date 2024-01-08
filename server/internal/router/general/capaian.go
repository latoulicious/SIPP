package general

import (
	"github.com/gofiber/fiber/v2"
)

func SetupCapaianRoutes(router fiber.Router) {
	capaian := router.Group("/capaian")

	capaian.Get("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /capaian")
	})

	capaian.Get("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /capaian/:id")
	})

	capaian.Post("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("POST /capaian")
	})

	capaian.Put("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("PUT /capaian/:id")
	})

	capaian.Delete("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("DELETE /capaian/:id")
	})
}
