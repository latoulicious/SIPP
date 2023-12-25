package core

import (
	"github.com/gofiber/fiber/v2"
)

func SetupTahunRoutes(router fiber.Router) {
	tahun := router.Group("/tahun")

	tahun.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("GET /tahun")
	})

	tahun.Post("/", func(c *fiber.Ctx) error {

		return c.SendString("POST /tahun")
	})

	tahun.Get("/:id", func(c *fiber.Ctx) error {

		return c.SendString("GET /tahun/:id")
	})

	tahun.Put("/:id", func(c *fiber.Ctx) error {

		return c.SendString("PUT /tahun/:id")
	})

	tahun.Delete("/:id", func(c *fiber.Ctx) error {

		return c.SendString("DELETE /tahun/:id")
	})

}
