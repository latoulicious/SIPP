package general

import (
	"github.com/gofiber/fiber/v2"
)

func SetupSoalRoutes(router fiber.Router) {
	soal := router.Group("/soal")

	soal.Get("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /soal")
	})

	soal.Get("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /soal/:id")
	})

	soal.Post("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("POST /soal")
	})

	soal.Put("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("PUT /soal/:id")
	})

	soal.Delete("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("DELETE /soal/:id")
	})
}
