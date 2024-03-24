package general

import (
	"github.com/gofiber/fiber/v2"
)

func SetupIndikatorRoutes(router fiber.Router) {
	indikator := router.Group("/indikator")

	indikator.Get("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /indikator")
	})

	indikator.Get("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /indikator/:id")
	})

	indikator.Post("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("POST /indikator")
	})

	indikator.Put("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("PUT /indikator/:id")
	})

	indikator.Delete("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("DELETE /indikator/:id")
	})
}
