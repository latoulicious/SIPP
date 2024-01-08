package general

import (
	"github.com/gofiber/fiber/v2"
)

func SetupBankSoalRoutes(router fiber.Router) {
	banksoal := router.Group("/banksoal")

	banksoal.Get("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /banksoal")
	})

	banksoal.Get("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /banksoal/:id")
	})

	banksoal.Post("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("POST /banksoal")
	})

	banksoal.Put("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("PUT /banksoal/:id")
	})

	banksoal.Delete("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("DELETE /banksoal/:id")
	})
}
