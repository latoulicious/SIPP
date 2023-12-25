package core

import (
	"github.com/gofiber/fiber/v2"
)

func SetupKelasRoutes(router fiber.Router) {
	kelas := router.Group("/kelas")

	kelas.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("GET /kelas")
	})

	kelas.Post("/", func(c *fiber.Ctx) error {

		return c.SendString("POST /kelas")
	})

	kelas.Get("/:id", func(c *fiber.Ctx) error {

		return c.SendString("GET /kelas/:id")
	})

	kelas.Put("/:id", func(c *fiber.Ctx) error {

		return c.SendString("PUT /kelas/:id")
	})

	kelas.Delete("/:id", func(c *fiber.Ctx) error {

		return c.SendString("DELETE /kelas/:id")
	})
}
