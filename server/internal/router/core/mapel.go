package core

import (
	"github.com/gofiber/fiber/v2"
)

func SetupMapelRoutes(router fiber.Router) {
	mapel := router.Group("/mapel")

	mapel.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("GET /mapel")
	})

	mapel.Get("/:id", func(c *fiber.Ctx) error {

		return c.SendString("GET /mapel/:id")
	})

	mapel.Post("/", func(c *fiber.Ctx) error {

		return c.SendString("POST /mapel")
	})

	mapel.Put("/:id", func(c *fiber.Ctx) error {

		return c.SendString("PUT /mapel/:id")
	})

	mapel.Delete("/:id", func(c *fiber.Ctx) error {

		return c.SendString("DELETE /mapel/:id")
	})

}
