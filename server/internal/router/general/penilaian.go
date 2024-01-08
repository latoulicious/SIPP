package general

import (
	"github.com/gofiber/fiber/v2"
)

func SetupPenilaianRoutes(router fiber.Router) {
	penilaian := router.Group("/penilaian")

	penilaian.Get("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /penilaian")
	})

	penilaian.Get("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /penilaian/:id")
	})

	penilaian.Post("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("POST /penilaian")
	})

	penilaian.Put("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("PUT /penilaian/:id")
	})

	penilaian.Delete("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("DELETE /penilaian/:id")
	})
}
