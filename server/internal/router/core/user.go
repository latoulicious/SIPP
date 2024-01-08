package core

import (
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")

	user.Get("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /user")
	})

	user.Get("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /user/:id")
	})

	user.Post("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("POST /user")
	})

	user.Put("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("PUT /user/:id")
	})

	user.Delete("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("DELETE /user/:id")
	})
}
