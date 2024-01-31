package core

import (
	"github.com/gofiber/fiber/v2"
)

func SetupJurusanRoutes(router fiber.Router) {
	jurusan := router.Group("/jurusan")

	jurusan.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("GET /jurusan")
	})

	jurusan.Get("/:id", func(c *fiber.Ctx) error {

		return c.SendString("GET /jurusan/:id")
	})

	jurusan.Post("/", func(c *fiber.Ctx) error {

		return c.SendString("POST /jurusan")
	})

	jurusan.Put("/:id", func(c *fiber.Ctx) error {

		return c.SendString("PUT /jurusan/:id")
	})

	jurusan.Delete("/:id", func(c *fiber.Ctx) error {

		return c.SendString("DELETE /jurusan/:id")
	})
}
